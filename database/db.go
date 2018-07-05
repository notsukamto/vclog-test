package database

import (
	"database/sql"
	"fmt"
	"os"
)

func envOrDefault(key string, def string) string {
	v := os.Getenv(key)
	if v == "" {
		v = def
	}
	return v
}

func dbURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		envOrDefault("DBUSER", "testing"),
		envOrDefault("DBPASS", "testing"),
		envOrDefault("DBHOST", "0.0.0.0:5432"),
		envOrDefault("DBNAME", "testing"),
	)
}

func openDB() *sql.DB {
	db, err := sql.Open("postgres", dbURL())
	if err != nil {
		panic(err)
	}
	return db
}

// AddRegistrationData add the inputted registration data to the database Registration table
func AddRegistrationData(data *Registration) {
	db := openDB()
	store := NewRegistrationStore(db)

	err := store.Insert(data)
	if err != nil {
		panic(err)
	}
}
