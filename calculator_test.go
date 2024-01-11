package main

import (
        "testing"
	"github.com/shopspring/decimal"
	"fmt"
)

func TestPayerCount(t *testing.T) {
	// payerCount(payments map[string]decimal.Decimal) int
	payments := map[string]decimal.Decimal {
		"lisa": decimal.NewFromFloat(100.0),
		"ferg": decimal.NewFromFloat(300.0),
	}
	// Set up map of names to decimals
	expected := 2
	actual := payerCount(payments)
	// perform check for int count
	if actual != expected {
		t.Fatalf(fmt.Sprintf("Expected %v payers but got %v",  expected, actual))
	}
}
