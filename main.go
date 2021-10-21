package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {

	// connect to database
	conn, err := sql.Open("pgx", "")
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}
	defer conn.Close()
	log.Println("Connected to database...")

	// test connection
	err = conn.Ping()
	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot ping database: %v\n", err))
	}
	log.Println("Pinged database...")

	// retrieve all rows from users table
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error retrieving rows from user table: %v\n", err))
	}
}

func getAllRows(conn *sql.DB) error {
	rows, err := conn.Query("select id, first_name, last_name from users")
	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot query users table: %v\n", err))
	}
	defer rows.Close()

	var firstName, lastName string
	var id int

	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println("Row: ", id, firstName, lastName)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(fmt.Sprintf("Error scanning row in user table: %v\n", err))
	}

	fmt.Println("--------------------------------------------")

	return nil
}
