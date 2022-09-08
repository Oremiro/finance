package storage

import (
	"context"
	"finance/internal/model"
)

type (
	ITinkoffPostgres interface {
		GetAllTinkoffTransactions(ctx context.Context) ([]*model.BankTransaction, error)
	}
	ITinkoffWebDriver interface {
		DownloadBankStatement(ctx context.Context)
	}
	ITinkoffFile interface {
		GetDataFromBankStatement(ctx context.Context, fileBase64 string)
	}
)
