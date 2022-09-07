package tinkoff

import "finance/internal/model"

type (
	GetAllResponse struct {
		BankTransactions []*model.BankTransaction `json:"BankTransactions"`
	}
)
