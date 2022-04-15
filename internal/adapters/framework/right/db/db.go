package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
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

func (da Adapter) CloseDbConnection() {
	err := da.db.Close()
	if err != nil {
		fmt.Println("close db error")
	}
}

func (da Adapter) AddHistory(answer int32, operation string) error {
	querystring, args, err := sq.Insert("arith_history").Columns("date", "answer", "operation").Values(time.Now(), answer, operation).ToSql()
	if err != nil {
		return err
	}

	_, err = da.db.Exec(querystring, args)

	if err != nil {
		return err
	}

	return nil
}
