package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/caneel7/keynote-api/config"
)

func Open(ctx context.Context, cfg config.Database) (*sql.DB, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	db, err := sql.Open("pgx", cfg.Url())
	if err != nil {
		return nil, fmt.Errorf("error instantiating db: %w", err)
	}

	db.SetConnMaxIdleTime(cfg.ConnectionTimeout)
	db.SetMaxOpenConns(cfg.MaxOpenConnections)
	db.SetMaxIdleConns(cfg.MaxIdleConnections)

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := db.PingContext(pingCtx); err != nil {
		db.Close()
		return nil, fmt.Errorf("error establishing connection: %w", err)
	}

	return db, nil
}