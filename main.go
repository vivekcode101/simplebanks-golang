package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/vivekcode101/simplebanks-golang/api"
	db "github.com/vivekcode101/simplebanks-golang/db/sqlc"
	"github.com/vivekcode101/simplebanks-golang/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
