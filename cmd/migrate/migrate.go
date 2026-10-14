package main

import (
	"context"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/caneel7/keynote-api/config"
	"github.com/caneel7/keynote-api/database"
	"github.com/caneel7/keynote-api/migrate"
	embed "github.com/caneel7/keynote-api/migrate/migrations"
)

func main() {
	env,err := config.Load()

	if err != nil {
		log.Fatalf("Error loading config : %v",err)
	}

	db,err := database.Open(context.Background(),env.Database)

	if err != nil {
		log.Fatalf("error : %v",err)
	}

	if err := migrate.Run(context.Background(), db, embed.FS); err != nil {
		log.Fatal(err)
	}
}