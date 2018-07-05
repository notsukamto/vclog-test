//go:generate kallax gen
//kallax migrate -i ./models/ -o ./migrations/ -n initial_schema
//go:generate kallax migrate up -d ./migrations/ --dsn "vclog:w9QgaRDNDbkg2WsGli83Uoh2@vc-dev-db01.c1ugusbzuf2l.ap-southeast-1.rds.amazonaws.com:5432/vclog?sslmode=disable" -v 1530676411

package models

import (
	"time"

	"gopkg.in/src-d/go-kallax.v1"
)

// Registration struct
type Registration struct {
	kallax.Model
	ID                     kallax.ULID `pk:""`
	//ApplicationID          kallax.ULID
	//PhoneCountry           string
	//PhoneNumber            string
	//DeviceIdentifier       string
	//ManufacturerDeviceCode string
	SourceIP               string
	DateRegistered         time.Time
}

// Login struct
type Login struct {
	kallax.Model
	AccountID   kallax.ULID `pk:""`
	SourceIP    string
	DateCreated time.Time
}
