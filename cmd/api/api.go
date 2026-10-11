package main

import (
	"context"
	"log"

	"github.com/caneel7/keynote-api/config"
	"github.com/caneel7/keynote-api/database"
	"github.com/caneel7/keynote-api/internal/httpserver"
)

func main () {

	env,err := config.Load()

	if err != nil {
		log.Fatalf("Error loading config : %v",err)
	}

	db,err := database.Open(context.Background(),env.Database)

	if err != nil {
		log.Fatalf("Error connecting Db : %v",err)
	}
	defer db.Close()

	entClient := database.NewEntClient(db)
	
	server := httpserver.NewServer(env.HTTP,entClient)
	log.Printf("Server Start On port : %v",server.Addr)
	log.Fatal(server.ListenAndServe())
}