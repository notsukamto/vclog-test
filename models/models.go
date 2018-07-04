//go:generate kallax gen

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
