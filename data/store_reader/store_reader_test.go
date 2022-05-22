// Package defines a reader of store.
package store_reader

import (
	"strings"
	"testing"

	"github.com/tlacuilose/transaction-reporter/domain/entities"
)

// Test that transactions can be read from a store.
func TestReadTransactionFromStore(t *testing.T) {
	testcsv := `Id,Date,Transaction
0,7/11,+60.5
1,7/28,-10.3
2,8/2,-20.46
3,8/13,+10
`
	storeReader := New(entities.DecodeFromStrings)
	csvReader := strings.NewReader(testcsv)
	transactions, err := storeReader.ReadTransactionsFromCsv(csvReader)
	if err != nil {
		t.Fatal(err)
	}

	tr := transactions[0]

	if tr.Id != 0 || tr.Date != "7/11" || tr.Amount != 60.5 {
		t.Fatal("Transaction was not decoded correctly")
	}
}
