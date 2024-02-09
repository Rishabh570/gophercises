package db

import (
	"fmt"
	"time"

	bolt "go.etcd.io/bbolt"
)

// When this is set in another file in the same package, getting nil when I access it in this file
var Database *bolt.DB

func LoadDatabase() (*bolt.DB, error) {
	// Open a connection
	var err error
	Database, err = bolt.Open("tasksDatabase.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		fmt.Println("Some error occurred while opening db connection")
		return nil, err
	}

	Database.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("tasks"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	return Database, nil
}
