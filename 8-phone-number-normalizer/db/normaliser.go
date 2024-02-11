package db

import (
	"database/sql"
	"phoneNormaliser/utils"
)

func NormalisePhoneTable(db *sql.DB, phoneNumbers []string) {
	frequencyMap := make(map[string]int)

	// Get sanitized phone numbers
	sanitizedPhones := utils.RemoveUnwantedRunes(phoneNumbers)

	// Store frequencies in a map
	for _, sanitizedPhone := range sanitizedPhones {
		frequencyMap[sanitizedPhone]++
	}

	// Normalize entry in DB
	for _, phone := range phoneNumbers {
		sanitizedPhone := utils.RemoveUnwantedRunes([]string{phone})
		if frequencyMap[sanitizedPhone[0]] == 1 {

			// Only sanitize when original input is not the same as sanitized phone value
			if phone != sanitizedPhone[0] {
				// delete and insert sanitized version
				DeleteRow(db, phone)
				InsertPhone(db, sanitizedPhone)
			}
		} else {
			// delete duplicate
			DeleteRow(db, phone)

			// Decrease frequency by one
			frequencyMap[sanitizedPhone[0]]--
		}
	}
}
