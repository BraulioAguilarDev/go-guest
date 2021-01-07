package database

import (
	"github.com/brauliodev29/go-guest/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Init func
func Init() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", config.Config.DatabaseURI)
	if err != nil {
		return nil, err
	}

	return db, nil
}
