package main

import (
	"context"
	"log"
	"os"

	"github.com/caneel7/keynote-api/config"
	"github.com/caneel7/keynote-api/database"
	"github.com/caneel7/keynote-api/internal/common/tokens"
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

	pem, err := os.ReadFile(env.JWT.PrivateKeyFile)
	key, err := tokens.ParsePrivateKeyPEM(pem)

	signer, err := tokens.NewSigner(
		env.JWT.Issuer,
		env.JWT.Audience,
		env.JWT.RefreshTokenTTL,
		env.JWT.AccessTTL,
		key,
	)
	
	server := httpserver.NewServer(env.HTTP,entClient,signer)
	log.Printf("Server Start On port %v",server.Addr)
	log.Printf("For OpenApi doc go to %v", server.Addr + "/docs")
	log.Fatal(server.ListenAndServe())
}