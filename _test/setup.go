package _test

import "go-todo/database"

func Setup() {
	const TEST_DATABASE_CONNECTION_STRING = "postgresql://go:password123@localhost:5433/go_db_test"

	database.ConnectToDB(TEST_DATABASE_CONNECTION_STRING)
	database.Migrate()
}
