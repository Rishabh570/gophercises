package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	pg "phoneNormaliser/db"

	_ "github.com/lib/pq"
)

// Update creds before running
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "phone-normaliser"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Must ping to confirm connection health, sql.Open doesn't validate connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Create a new phone numbers table
	pg.CreateTable(db)

	// Insert some rows
	phones := []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892"}
	pg.InsertPhone(db, phones)

	// Fetch all the numbers from the database
	phoneNumbers, err := pg.GetPhoneNumbers(db)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Normalize rows
	pg.NormalisePhoneTable(db, phoneNumbers)
}
