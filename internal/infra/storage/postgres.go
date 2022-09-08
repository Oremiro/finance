package storage

import (
	"context"
	"finance/internal/model"
	"finance/pkg/postgres"
	"log"
)

type Postgres struct {
	Context *postgres.Context
}

func (p *Postgres) StoreTinkoffTransactionsFromBankStatements(ctx context.Context, transactions []*model.BankTransaction) {
	ib := p.Context.Builder.
		Insert("bank_transactions").
		Columns("operation_date", "payment_date", "card_number", "status", "operation", "currency", "payment", "payment_currency", "cashback", "category", "description")
	for _, t := range transactions {
		ib = ib.Values(t.OperationDate, t.PaymentDate, t.CardNumber, t.Status, t.Operation, t.Currency, t.Payment, t.PaymentCurrency, t.Cashback, t.Category, t.Description)
	}
	sql, args, err := ib.ToSql()
	if err != nil {
		log.Fatal(err)
	}
	_, err = p.Context.Pool.Exec(ctx, sql, args...)
	if err != nil {
		log.Fatal(err)
	}

}

func (p *Postgres) GetAllTinkoffTransactions(ctx context.Context) ([]*model.BankTransaction, error) {
	// TODO query is future paged request
	sql, _, err := p.Context.Builder.
		Select("id", "operation_date", "payment_date", "status", "operation").
		From("bank_transactions").
		OrderBy("id").
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := p.Context.Pool.Query(ctx, sql)
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

func NewPostgres(context *postgres.Context) *Postgres {
	return &Postgres{Context: context}
}
