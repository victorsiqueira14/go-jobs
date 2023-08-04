package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	// initialize DB
	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("Error initializing DB: %v", err)
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// initialize logger
	logger = NewLogger(p)
	return logger
}
