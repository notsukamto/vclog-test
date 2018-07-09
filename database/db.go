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
		envOrDefault("DBUSER", "vclog"),
		envOrDefault("DBPASS", "w9QgaRDNDbkg2WsGli83Uoh2"),
		envOrDefault("DBHOST", "vc-dev-db01.c1ugusbzuf2l.ap-southeast-1.rds.amazonaws.com"),
		envOrDefault("DBNAME", "vclog"),
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

// AddLoginData add the inputted login data to the database Login table
func AddLoginData(data *Login) {
	db := openDB()
	store := NewLoginStore(db)

	err := store.Insert(data)
	if err != nil {
		panic(err)
	}
}
