package main

import (
	"fmt"
	"log"

	"github.com/themy4a4o-hub/mycrypto/internal/config"
	"github.com/themy4a4o-hub/mycrypto/internal/db"
)

func main() {
	cfg := config.NewConfigMust()
	db, err := db.NewPostgres(*cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("DB connected")
	defer db.Close()
}
