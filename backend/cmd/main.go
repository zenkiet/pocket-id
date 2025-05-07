package main

import (
	"log"

	"github.com/pocket-id/pocket-id/backend/internal/bootstrap"
	_ "time/tzdata"
)

// @title Pocket ID API
// @version 1.0
// @description.markdown

func main() {
	err := bootstrap.Bootstrap()
	if err != nil {
		log.Fatal(err.Error())
	}
}
