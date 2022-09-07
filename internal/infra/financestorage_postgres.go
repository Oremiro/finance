package infra

import "finance/pkg/postgres"

type (
	FinanceStorage struct {
		Context *postgres.Context
	}
)

func NewFinanceStorage(context *postgres.Context) *FinanceStorage {
	return &FinanceStorage{Context: context}
}
