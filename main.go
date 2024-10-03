package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/omsatish/simplebank/api"
	db "github.com/omsatish/simplebank/db/sqlc"
	"github.com/omsatish/simplebank/db/util"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("Cannot load config file")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
