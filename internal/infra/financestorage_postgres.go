package infra

import (
	"context"
	"finance/internal/model"
	"finance/pkg/postgres"
)

type (
	ITinkoffStorage interface {
		GetAllTinkoffTransactions(ctx context.Context) ([]*model.BankTransaction, error)
	}
	FinanceStorage struct {
		Context *postgres.Context
	}
)

func (f *FinanceStorage) GetAllTinkoffTransactions(ctx context.Context) ([]*model.BankTransaction, error) {
	// TODO query is future paged request
	sql, _, err := f.Context.Builder.
		Select("id", "operation_date", "payment_date", "status", "operation").
		From("bank_transactions").
		OrderBy("id").
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := f.Context.Pool.Query(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	items := make([]*model.BankTransaction, 0)
	for rows.Next() {
		m := &model.BankTransaction{}

		err = rows.Scan(&m.ID, &m.OperationDate, &m.PaymentDate, &m.Status, &m.Operation)

		if err != nil {
			return nil, err
		}

		items = append(items, m)
	}

	return items, nil
}

func NewFinanceStorage(context *postgres.Context) *FinanceStorage {
	return &FinanceStorage{Context: context}
}
