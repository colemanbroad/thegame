package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // Import SQLite driver
)

func main() {
	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", "./example.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// Create a new table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY,
        name TEXT,
        age INTEGER
    )`)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	// Insert a new record into the table
	_, err = db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", "Alice", 30)
	if err != nil {
		fmt.Println("Error inserting record:", err)
		return
	}

	// Query the table
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println("Error querying database:", err)
		return
	}
	defer rows.Close()

	// Iterate over the result set
	for rows.Next() {
		var id, age int
		var name string
		if err := rows.Scan(&id, &name, &age); err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
		return
	}
}
