package util

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func CreateDB() (*sql.DB) {

	file, err := os.Create("db_test.db")
	if err != nil {
		log.Fatal(err)
	}

	file.Close()

	db, err := sql.Open("sqlite3", "db_test.db")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func DropDB(db *sql.DB) {
	err := os.Remove("db_test.db")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func QueryExec(db *sql.DB, command string) {
	query, err := db.Prepare(command)
	if err != nil {
		log.Fatal(err)
	}

	_, err = query.Exec()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Query executed successfully!")
}

func ClearTable(db *sql.DB, name string) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.Exec("delete from " + name)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
