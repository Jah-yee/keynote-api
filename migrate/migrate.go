package migrate

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"io/fs"
	"log"
	"sort"
)

const advisoryLockID int64 = 727274

const createTableSQL = `
CREATE TABLE IF NOT EXISTS schema_migrations (
	filename   text        PRIMARY KEY,
	checksum   text        NOT NULL,
	applied_at timestamptz NOT NULL DEFAULT now()
)`

// Run applies, in filename (= timestamp) order, every *.up.sql file in fsys
func Run(ctx context.Context, db *sql.DB, fsys fs.FS) error {
	conn, err := db.Conn(ctx)
	if err != nil {
		return fmt.Errorf("get connection: %w", err)
	}
	defer conn.Close()

	if _, err := conn.ExecContext(ctx, `SELECT pg_advisory_lock($1)`, advisoryLockID); err != nil {
		return fmt.Errorf("acquire lock: %w", err)
	}
	defer conn.ExecContext(context.Background(), `SELECT pg_advisory_unlock($1)`, advisoryLockID)

	if _, err := conn.ExecContext(ctx, createTableSQL); err != nil {
		return fmt.Errorf("create schema_migrations: %w", err)
	}

	applied, err := loadApplied(ctx, conn)
	if err != nil {
		return err
	}

	files, err := fs.Glob(fsys, "*.up.sql")
	if err != nil {
		return fmt.Errorf("list migrations: %w", err)
	}
	sort.Strings(files)

	onDisk := make(map[string]bool, len(files))
	pending := 0

	for _, name := range files {
		onDisk[name] = true

		content, err := fs.ReadFile(fsys, name)
		if err != nil {
			return fmt.Errorf("read %s: %w", name, err)
		}
		sum := checksum(content)

		if prev, ok := applied[name]; ok {
			if prev != sum {
				return fmt.Errorf("migration %s was modified after being applied (recorded %s, now %s)", name, prev, sum)
			}
			continue
		}

		if err := applyOne(ctx, conn, name, string(content), sum); err != nil {
			return fmt.Errorf("apply %s: %w", name, err)
		}
		log.Printf("applied %s", name)
		pending++
	}

	for name := range applied {
		if !onDisk[name] {
			log.Printf("warning: %s is recorded as applied but missing from migrations", name)
		}
	}

	if pending == 0 {
		log.Println("no pending migrations")
	}
	return nil
}

func applyOne(ctx context.Context, conn *sql.Conn, name, query, sum string) error {
	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, query); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx,
		`INSERT INTO schema_migrations (filename, checksum) VALUES ($1, $2)`, name, sum); err != nil {
		return err
	}
	return tx.Commit()
}

func loadApplied(ctx context.Context, conn *sql.Conn) (map[string]string, error) {
	rows, err := conn.QueryContext(ctx, `SELECT filename, checksum FROM schema_migrations`)
	if err != nil {
		return nil, fmt.Errorf("load applied migrations: %w", err)
	}
	defer rows.Close()

	applied := make(map[string]string)
	for rows.Next() {
		var name, sum string
		if err := rows.Scan(&name, &sum); err != nil {
			return nil, err
		}
		applied[name] = sum
	}
	return applied, rows.Err()
}

func checksum(b []byte) string {
	h := sha256.Sum256(b)
	return hex.EncodeToString(h[:])
}