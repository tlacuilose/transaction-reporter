// Package defines the entities of the project.
package entities

import (
	"fmt"
	"testing"
)

// Test that a new transaction can be created.
func TestNewTransaction(t *testing.T) {
	var id uint64 = 0
	date := "5/12"
	amount := 11.2
	transaction := Transaction{id, date, amount}
	if transaction.Id != id || transaction.Date != date || transaction.Amount != amount {
		t.Fatalf("Failed to create a new transaction")
	}
}

// Test that a transaction can be checked if it is credit.
func TestCreditTransaction(t *testing.T) {
	transaction := Transaction{0, "5/21", 100.0}
	if !transaction.IsCredit() {
		t.Fatal("Failed to check if transaction is credit.")
	}
}

// Test that a transaction can be checked if it is debit.
func TestDebitTransaction(t *testing.T) {
	transaction := Transaction{0, "5/21", -100.0}
	if !transaction.IsDebit() {
		t.Fatal("Failed to check if transaction is debit.")
	}
}

// Test that a transaction acan be decoded from an array of strings.
func TestDecodeTransaction(t *testing.T) {
	var tests = [][]string{
		{"1", "2/12", "-23.4"},
		{"0", "2/12", "23.4"},
		{"1", "2/12", "0"},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("Converting %v to Transaction", test)
		t.Run(testname, func(t *testing.T) {
			_, err := DecodeFromStrings(test)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
