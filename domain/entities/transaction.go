// Package defines the entities of the project.
package entities

import (
	"errors"
	"strconv"
)

// Transaction represents can be debit or credit.
type Transaction struct {
	Id     uint64
	Date   string
	Amount float64
}

// IsCredit checks if the transaction is credit.
func (t Transaction) IsCredit() bool {
	return t.Amount >= 0
}

// IsDebit checks if the transaction is credit.
func (t Transaction) IsDebit() bool {
	return t.Amount < 0
}

// DecodeFromStrings takes an array of strings and decodes it into a transaction.
func DecodeFromStrings(encoded []string) (Transaction, error) {
	id, err := strconv.ParseUint(encoded[0], 10, 64)
	if err != nil {
		return Transaction{}, errors.New("Failed to decode transaction")
	}

	date := encoded[1]

	amount, err := strconv.ParseFloat(encoded[2], 64)
	if err != nil {
		return Transaction{}, errors.New("Failed to decode transaction")
	}

	return Transaction{id, date, amount}, nil
}
