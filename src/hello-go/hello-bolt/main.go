package main

import (
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

func main() {
	fmt.Println("hi I will learn bolt database")

	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	// db, err := bolt.Open("my.db", 0600, nil)
	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_ = func() error {
		// Start a writable transaction.
		tx, err := db.Begin(true)
		if err != nil {
			return err
		}
		defer tx.Rollback()

		// Use the transaction...
		bkt, err := tx.CreateBucket([]byte("MyBucket"))
		if err != nil {
			return err
		}
		_ = bkt

		// Commit the transaction and check for error.
		if err := tx.Commit(); err != nil {
			return err
		}
		return nil
	}()

	fmt.Println("over!")
}
