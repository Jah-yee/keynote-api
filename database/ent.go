package database

import (
	"database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"

	"github.com/caneel7/keynote-api/internal/ent/gen"
)

func NewEntClient(db *sql.DB) *gen.Client {

	drv := entsql.OpenDB(dialect.Postgres,db)
	return gen.NewClient(gen.Driver(drv))

}
