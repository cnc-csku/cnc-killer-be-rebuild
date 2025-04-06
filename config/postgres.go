package config

import (
	"context"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func ConnectDatabase(cfg Config, ctx context.Context) *sqlx.DB {
	pgURI := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUsername, cfg.DBPassword, cfg.DBDatabase, cfg.DBSSLMode)

	db, err := sqlx.ConnectContext(ctx, "pgx", pgURI)
	if err != nil {
		log.Fatalf("error! : %s", err)
	}

	return db
}
