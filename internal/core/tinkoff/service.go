package tinkoff

import (
	"context"
	"finance/internal/infra/storage"
	"finance/internal/model"
)

type ITinkoffService interface {
	GetAll(ctx context.Context, query GetAllQuery) (*GetAllResponse, error)
	UpdateBankTransactions(ctx context.Context, command *UpdateBankTransactionsCommand) error
}

type Service struct {
	pg        storage.ITinkoffPostgres
	webdriver storage.ITinkoffWebDriver
	file      storage.ITinkoffFile
}

func (s *Service) UpdateBankTransactions(command *UpdateBankTransactionsCommand) error {
	ctx := context.Background()
	statements, err := s.file.GetDataFromBankStatement(ctx, command.FileBase64)
	if err != nil {
		return err
	}
	// map statement and transaction
	transactions := make([]*model.BankTransaction, 0, len(statements))
	for _, statement := range statements {
		transaction := &model.BankTransaction{
			Bank:            &model.Bank{},
			OperationDate:   statement.OperationDate,
			PaymentDate:     statement.PaymentDate,
			CardNumber:      statement.CardNumber,
			Status:          statement.Status,
			Operation:       statement.Operation,
			Currency:        statement.Currency,
			Payment:         statement.Payment,
			PaymentCurrency: statement.PaymentCurrency,
			Cashback:        statement.Cashback,
			Category:        statement.Category,
			Description:     statement.Description,
		}
		transactions = append(transactions, transaction)
	}

	s.pg.StoreTinkoffTransactionsFromBankStatements(ctx, transactions)
	ctx.Done()
	return nil
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
