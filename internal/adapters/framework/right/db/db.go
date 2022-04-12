package db

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, datasource string) (*Adapter, error) {
	db, err := sql.Open(driverName, datasource)
	if err != nil {
		log.Fatal("db connection error")
	}
	ping := db.Ping()
	if ping != nil {
		log.Fatal("ping fail")
	}
	return &Adapter{db: db}, nil
}
