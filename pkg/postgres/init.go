package postgres

import (
	"SberApi/config"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	queryCreateTable = `CREATE TABLE IF NOT EXISTS task (
		id VARCHAR(100),
		title VARCHAR(100),
		description VARCHAR(100),
		data DATE,
		checking_completion BOOLEAN NOT NULL
	  );`
)

func InitPsqlDB(ctx context.Context, cfg *config.Config) (*sqlx.DB, error) {
	connectionURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DBName,
		cfg.Postgres.SSLMode,
	)

	database, err := sqlx.Open("postgres", connectionURL)
	if err != nil {
		return nil, err
	}

	if err = database.Ping(); err != nil {
		return nil, err
	}

	database.MustExec(queryCreateTable)

	return database, nil
}