package main

import (
	"log"

	"github.com/caneel7/keynote-api/internal/httpserver"
	"github.com/caneel7/keynote-api/config"
)

func main () {

	env,err := config.Load()

	if err != nil {
		log.Fatalf("Error loading config : %v",err)
	}
	
	server := httpserver.NewServer(env.HTTP)
	log.Printf("Server Start On port : %v",server.Addr)
	log.Fatal(server.ListenAndServe())
}