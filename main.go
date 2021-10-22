package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {

	// connect to database
	conn, err := sql.Open("", "")
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

	// insert a row into the users table
	query := `insert into users (first_name, last_name) values ($1, $2)`
	_, err = conn.Exec(query, "Paul", "Hanover")
	if err != nil {
		log.Fatal(fmt.Sprintf("Error inserting rows into user table: %v\n", err))
	}

	log.Println("Row inserted into users table...")

	// retrieve all rows from users table to test insert
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error retrieving rows from user table: %v\n", err))
	}

	// update a row in users table
	updateStmt := `update users set first_name = $1 where first_name = $2`
	_, err = conn.Exec(updateStmt, "Mary Ellen", "Mary")
	if err != nil {
		log.Fatal(fmt.Sprintf("Error updating row in user table: %v\n", err))
	}

	log.Println("Update row  uin sers table...")

	// retrieve all rows from users table to test update
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error retrieving rows from user table: %v\n", err))
	}

	// get one row by id
	selectOneRowQuery := `select id, first_name, last_name from users where id = $1`

	var firstName, lastName string
	var id int

	row := conn.QueryRow(selectOneRowQuery, 1)
	err = row.Scan(&id, &firstName, &lastName)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("QueryRow returns", id, firstName, lastName)

	// delete a row
	deleteStmt := `delete from users where id = $1`
	_, err = conn.Exec(deleteStmt, 1)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Deleted a row!")

	// get rows from table again
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
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
