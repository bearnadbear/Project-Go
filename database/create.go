package database

import (
	"os"

	"gorm.io/gorm"
)

func CreateTable(db *gorm.DB) error {
	// readfile to 'db.sql'
	read, err := os.ReadFile("database/db.sql")
	if err != nil {
		return err
	}

	// Casting []byte to string
	newFile := string(read)

	// Execute file 'db.sql'
	db.Exec(newFile)

	return nil
}
