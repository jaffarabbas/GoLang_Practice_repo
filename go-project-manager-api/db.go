package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type MySqlStorage struct {
	db *sql.DB
}

func NewMySqlStorage(cfg mysql.Config) *MySqlStorage {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	log.Println("Connected to database")

	return &MySqlStorage{db: db}
}

func (s *MySqlStorage) init() (*sql.DB, error) {
	return s.db, nil
}
