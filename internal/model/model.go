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

const (
	TINKOFF Bank = iota
)

type (
	Status          int
	Currency        int
	Bank            int
	Category        string
	BankTransaction struct {
		ID                     uint64    `json:"ID"`                     // ID
		Bank                   Bank      `json:"Bank"`                   // Bank //TODO foreign key
		OperationDate          time.Time `json:"OperationDate"`          // Date of operation
		PaymentDate            time.Time `json:"PaymentDate"`            // Payment date
		CardNumber             string    `json:"CardNumber"`             // Card number
		Status                 Status    `json:"Status"`                 // Status
		Operation              float64   `json:"Operation"`              // Operation amount
		Currency               Currency  `json:"Currency"`               // Transaction currency
		Payment                float64   `json:"Payment"`                // Amount of payment
		PaymentCurrency        Currency  `json:"PaymentCurrency"`        // Payment currency
		Cashback               float64   `json:"Cashback"`               // Cashback
		Category               Category  `json:"Category"`               // Category
		MCC                    int64     `json:"MCC"`                    // MCC
		Description            string    `json:"Description"`            // Description
		Bonuses                float64   `json:"Bonuses"`                // Bonuses (including cashback)
		InvestmentBankRounding float64   `json:"InvestmentBankRounding"` // Rounding per investment bank
		Rounding               float64   `json:"Rounding"`               // The amount of the operation with rounding
	}
)
