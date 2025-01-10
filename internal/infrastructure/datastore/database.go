package datastore

import (
	"net/url"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDatabase(databaseURL string) (db *sqlx.DB, err error) {
	parseDBUrl, _ := url.Parse(databaseURL)
	db, err = sqlx.Open(parseDBUrl.Scheme, databaseURL)
	if err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		return
	}

	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxIdleConns(0)
	db.SetMaxOpenConns(5)

	return
}
