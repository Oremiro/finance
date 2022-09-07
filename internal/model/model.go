// Package model Can be split in several packages
package model

import "time"

const (
	OK     Status = iota
	FAILED Status = iota
)

const (
	RUB Currency = iota
)

type (
	Status          int
	Currency        int
	Category        string
	BankTransaction struct {
		OperationDate          time.Time // Date of operation
		PaymentDate            time.Time // Payment date
		CardNumber             string    // Card number
		Status                 Status    // Status
		Operation              int64     // Operation amount
		Currency               Currency  // Transaction currency
		Payment                int64     // Amount of payment
		PaymentCurrency        Currency  // Payment currency
		Cashback               int64     // Cashback
		Category               Category  // Category
		MCC                    int64     // MCC
		Description            string    // Description
		Bonuses                int64     // Bonuses (including cashback)
		InvestmentBankRounding int64     // Rounding per investment bank
		Rounding               int64     // The amount of the operation with rounding
	}
)
