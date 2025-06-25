package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // import MySQL driver
)

func main() {
	// Connect to MySQL (username:password@tcp(host:port)/dbname)
	db, err := sql.Open("mysql", "root:Mano@2216@tcp(127.0.0.1:3306)/student")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ping to verify connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to DB!")

	// Query data
	rows, err := db.Query("SELECT id, name FROM employee")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Loop through rows
	for rows.Next() {
		var id int
		var name string

		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
