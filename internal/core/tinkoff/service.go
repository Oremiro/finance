package tinkoff

import (
	"context"
	"finance/internal/infra"
)

type ITinkoffService interface {
	GetAll(ctx context.Context, query GetAllQuery) (*GetAllResponse, error)
}

type Service struct {
	storage *infra.FinanceStorage
}

func (s *Service) GetAll(ctx context.Context, query GetAllQuery) (*GetAllResponse, error) {
	items, err := s.storage.GetAllTinkoffTransactions(ctx)
	if err != nil {
		return nil, err
	}
	response := &GetAllResponse{
		BankTransactions: items,
	}
	return response, nil
}

func NewService(storage *infra.FinanceStorage) *Service {
	return &Service{storage: storage}
}
