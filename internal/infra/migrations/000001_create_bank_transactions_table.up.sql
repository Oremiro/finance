create table if not exists bank_transactions
(
    id                       serial primary key,
    operation_date           timestamp,
    payment_date             timestamp,
    card_number              varchar(50),
    status                   int,
    operation                float,
    currency                 int,
    payment                  float,
    payment_currency         int,
    cashback                 float,
    category                 varchar(50),
    mcc                      int,
    description              varchar(256),
    bonuses                  float,
    investment_bank_rounding float,
    rounding                 float
);