package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"kodboris/api"
	db "kodboris/db/sqlc"
	"kodboris/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load env: %w", err)
	}
	fmt.Printf("Database driver: %v\n", config.DbDriver)

	conn, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatal("Error connecting to db:", err)
	}
	// Fail startup, if connection to the db can not be established
	err = conn.Ping() // Check if the connection is valid
	if err != nil {
		log.Fatalf("Application can not start, Error connecting to db: %v", err) // Fatal log and terminate
	}

	store := db.NewStore(conn)

	server := &api.Server{}
	address := config.ServerAddress
	err = server.Start(address, store)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}

}
