package storage

import (
	"context"
	"finance/internal/model"
)

type (
	ITinkoffPostgres interface {
		GetAllTinkoffTransactions(ctx context.Context) ([]*model.BankTransaction, error)
		StoreTinkoffTransactionsFromBankStatements(ctx context.Context, transactions []*model.BankTransaction)
	}
	ITinkoffWebDriver interface {
		DownloadBankStatement(ctx context.Context)
	}
	ITinkoffFile interface {
		GetDataFromBankStatement(ctx context.Context, fileBase64 string) ([]*model.TinkoffBankStatement, error)
	}
)
