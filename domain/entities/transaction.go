// Package defines the entities of the project.
package entities

import (
	"errors"
	"strconv"
	"strings"
)

// Transaction represents can be debit or credit.
type Transaction struct {
	Id     uint64
	Date   string
	Amount float64
}

// Check date format is valid dd/mm
func (t Transaction) checkDateFormat() (uint8, uint8, error) {
	dateFormat := strings.Split(t.Date, "/")
	if len(dateFormat) != 2 {
		return 0, 0, errors.New("Failed to decode transaction, data format is mm/dd")
	}

	month, err := strconv.ParseUint(dateFormat[0], 10, 8)
	if err != nil || month > 12 {
		return 0, 0, errors.New("Failed to decode transaction, data format is mm/dd")
	}

	day, err := strconv.ParseUint(dateFormat[1], 10, 8)
	if err != nil || day > 31 {
		return 0, 0, errors.New("Failed to decode transaction, data format is mm/dd")
	}
	return uint8(month), uint8(day), nil
}

// IsCredit checks if the transaction is credit.
func (t Transaction) IsCredit() bool {
	return t.Amount >= 0
}

// IsDebit checks if the transaction is credit.
func (t Transaction) IsDebit() bool {
	return t.Amount < 0
}

// GetMonth gets the month of the transaction.
func (t Transaction) GetMonth() (uint8, error) {
	month, _, err := t.checkDateFormat()
	return month, err
}

// GetDay gets the day of the transaction.
func (t Transaction) GetDay() (uint8, error) {
	_, day, err := t.checkDateFormat()
	return day, err
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

	transaction := Transaction{id, date, amount}
	if _, _, err = transaction.checkDateFormat(); err != nil {
		return Transaction{}, err
	}

	return Transaction{id, date, amount}, nil
}
