package config

import (
	"os"

	"github.com/victorsiqueira14/go-jobs/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	// check if DB exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Infof("Creating not founded DB at %s", dbPath)
		// create DB
		err := os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("Error creating DB: %v", err)
			return nil, err
		}
		// create file
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		// close file
		file.Close()
	}

	// create DB and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error connecting to DB: %v", err)
		return nil, err
	}

	// migrate schemas
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Error migrating DB: %v", err)
		return nil, err
	}
	// return the DB
	return db, nil
}
