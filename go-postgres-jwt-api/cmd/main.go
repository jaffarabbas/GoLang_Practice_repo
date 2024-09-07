package main

import (
	"database/sql"
	"go-jwt-postgres-api/cmd/api"
	"go-jwt-postgres-api/config"
	"go-jwt-postgres-api/db"
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySqlStorage(mysql.Config{
		User:      config.Envs.DBUser,
		Passwd:    config.Envs.DBPassword,
		Addr:      config.Envs.DBAddress,
		DBName:    config.Envs.DBName,
		Net:       "tcp",
		ParseTime: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	initStorage(db)
	server := api.NewApiServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DB: Successfully connected!")
}
