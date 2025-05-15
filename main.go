package main

import (
	"database/sql"
	"log"

	"github.com/gajanan-deshmukh/api"
	db "github.com/gajanan-deshmukh/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@postgres12:5432/simple_bank?sslmode=disable"
	address  = "0.0.0.0:8081"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connected to datbase")
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(address)
	if err != nil {
		log.Fatal("failed to start server ")
	}

}
