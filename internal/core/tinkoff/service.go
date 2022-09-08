package tinkoff

import (
	"context"
	"finance/internal/infra/storage"
)

type ITinkoffService interface {
	GetAll(ctx context.Context, query GetAllQuery) (*GetAllResponse, error)
	UpdateBankTransactions(ctx context.Context) error
}

type Service struct {
	pg        storage.ITinkoffPostgres
	webdriver storage.ITinkoffWebDriver
	file      storage.ITinkoffFile
}

func (s *Service) UpdateBankTransactions(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetAll(ctx context.Context, query GetAllQuery) (*GetAllResponse, error) {
	items, err := s.pg.GetAllTinkoffTransactions(ctx)
	if err != nil {
		return nil, err
	}
	response := &GetAllResponse{
		BankTransactions: items,
	}
	return response, nil
}

func NewService(pg storage.ITinkoffPostgres, webdriver storage.ITinkoffWebDriver, file storage.ITinkoffFile) *Service {
	return &Service{pg: pg, webdriver: webdriver, file: file}
}
