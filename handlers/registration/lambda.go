package main

import (
	"database/sql"
	"fmt"
	"os"

	"gopkg.in/src-d/go-kallax.v1"

	"github.com/notsukamto/vclog-test/models"
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

var schemas = []string{
	`CREATE TABLE login (
		account_id uuid NOT NULL PRIMARY KEY,
		source_ip text NOT NULL,
		date_created timestamptz NOT NULL
	)`,
	`CREATE TABLE registration (
		id uuid NOT NULL PRIMARY KEY,
		source_ip text NOT NULL,
		date_registered timestamptz NOT NULL
	)`,
}

var tables = []string{"login", "registration"}

func setupDB(db *sql.DB) *sql.DB {
	for _, s := range schemas {
		_, err := db.Exec(s)
		if err != nil {
			panic(err)
		}
	}

	return db
}

func teardownDB(db *sql.DB) {
	for _, t := range tables {
		_, err := db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", t))
		if err != nil {
			panic(err)
		}
	}

	if err := db.Close(); err != nil {
		panic(err)
	}
}

func makeRegistration() *models.Registration {
	return &models.Registration{
		ID:       kallax.NewULID(),
		SourceIP: "182.255.12.1",
	}
}

// AddRegistrationData add the inputted registration data to the database Registration table
func AddRegistrationData(data *models.Registration) {
	db := openDB()
	store := models.NewRegistrationStore(db)

	err := store.Insert(data)
	if err != nil {
		panic(err)
	}
}
