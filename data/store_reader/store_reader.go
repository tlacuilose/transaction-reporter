// Package defines a reader of store.
package store_reader

import (
	"encoding/csv"
	"errors"
	"io"

	"github.com/tlacuilose/transaction-reporter/domain/entities"
)

// StoreReader is used to read store.
type StoreReader struct {
	decoder func([]string) (entities.Transaction, error)
}

// New creates a new store reader with an decoder function.
func New(decoder func([]string) (entities.Transaction, error)) StoreReader {
	return StoreReader{decoder}
}

// ReadTransactions read transactions from a io.Reader
func (br StoreReader) ReadTransactionsFromCsv(reader io.Reader) ([]entities.Transaction, error) {
	transactions := make([]entities.Transaction, 0)
	r := csv.NewReader(reader)

	headers, err := r.Read()
	if err == io.EOF {
		return transactions, errors.New("Reader had no transactions")
	}
	if err != nil {
		return transactions, errors.New("Failed to get record from reader")
	}

	if headers[0] != "Id" || headers[1] != "Date" || headers[2] != "Transaction" {
		return transactions, errors.New("Badly formated headers")
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return make([]entities.Transaction, 0), errors.New("Failed to get record from reader")
		}

		t, err := br.decoder(record)
		if err != nil {
			return make([]entities.Transaction, 0), errors.New("Failed to decode transaction.")
		}

		transactions = append(transactions, t)
	}
	return transactions, nil
}
