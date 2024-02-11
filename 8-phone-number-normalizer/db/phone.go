package db

import (
	"database/sql"
	"fmt"
	"log"
)

func InsertPhone(db *sql.DB, phones []string) {
	for _, phone := range phones {
		_, err := db.Exec(`
			INSERT INTO phone (value)
			VALUES ($1)
		`, phone)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func GetPhoneNumbers(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT value FROM phone")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var phones []string

	// Iterate over the rows and print the values
	for rows.Next() {
		var phone string
		if err := rows.Scan(&phone); err != nil {
			log.Fatal(err)
			return nil, err
		}
		fmt.Println("Phone:", phone)
		phones = append(phones, phone)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return phones, nil
}

func CreateTable(db *sql.DB) error {
	// Create a new table
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS phone (
		id SERIAL PRIMARY KEY,
		value VARCHAR(20) UNIQUE NOT NULL
	)
	`)

	if err != nil {
		log.Fatal(err)
	}
	return err
}

func DeleteRow(db *sql.DB, phone string) error {
	_, err := db.Exec(`
			DELETE FROM phone WHERE value = ($1)
		`, phone)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
