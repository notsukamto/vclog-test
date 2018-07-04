package main

import(
)

type Response struct {
	Message string `json:"The data has been successfully inputted!"`
}

func insertRegistrationLog() (Response, error) {
	store := NewRegistrationStore
}
