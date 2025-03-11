package main

import "database/sql"

// This might actually be about databases before auth. (Because the database will generate the uuid? And have session logic?)

func temp_main() {
	connStr := "user=postgres dbname=test_db password=1234 host=localhost sslmode=disable"
	// The connection string can be tested: $ psql "...conn string..."
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Ping to test the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

// Why golang migrate? It's production tested and simple
// https://hn.algolia.com/?dateRange=all&page=0&prefix=true&query=golang-migrate&sort=byPopularity&type=all
