package database

import (
	"github.com/brauliodev29/go-guest/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Init func
func Init() (*gorm.DB, error) {
	db, err := gorm.Open(config.Config.Dialect, config.Config.DatabaseURI)
	if err != nil {
		return nil, err
	}

	return db, nil
}
