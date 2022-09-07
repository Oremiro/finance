package api

import "finance/pkg/postgres"

type (
	Config struct {
		DbOptions struct {
			ConnectionString *postgres.ConnectionString `json:"ConnectionString"`
		} `json:"DbOptions"`
		HttpServer struct{}
	}
)
