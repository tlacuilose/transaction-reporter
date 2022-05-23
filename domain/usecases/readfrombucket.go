// Package defines usecases
// Package defines usecases for project.
package usecases

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/tlacuilose/transaction-reporter/data/store_reader"
	"github.com/tlacuilose/transaction-reporter/domain/entities"
)

// ReadFromBucket reads transactions from csv in aws bucket.
func ReadFromBucket(bucket string, storeName string, sess *session.Session) ([]entities.Transaction, error) {
	fileName := "/tmp/" + storeName
	file, err := os.Create(fileName)
	if err != nil {
		return make([]entities.Transaction, 0), err
	}
	defer file.Close()

	downloader := s3manager.NewDownloader(sess)
	_, err = downloader.Download(
		file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(storeName),
		},
	)
	if err != nil {
		return make([]entities.Transaction, 0), err
	}

	storeReader := store_reader.New(entities.DecodeFromStrings)
	transactions, err := storeReader.ReadTransactionsFromCsv(file)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
