// Package defines usecases for project.
package usecases

import (
	"testing"
)

// Test that transactions can be read from a test store csv file.
func TestReadFromLocalTestFile(t *testing.T) {
	filename := "../../store/tests/001.csv"

	transactions, err := ReadFromLocalFile(filename)
	if err != nil {
		t.Fatal(err)
	}

	tr := transactions[0]
	if tr.Id != 0 || tr.Date != "7/15" || tr.Amount != 60.5 {
		t.Fatal("Failed to read file correctly")
	}
}
