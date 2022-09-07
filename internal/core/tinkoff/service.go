package tinkoff

import "finance/internal/infra"

type (
	Service struct {
		storage *infra.FinanceStorage
	}
)

func NewService(storage *infra.FinanceStorage) *Service {
	return &Service{storage: storage}
}
