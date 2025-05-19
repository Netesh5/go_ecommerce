package main

import (
	"log"
	"os"
)

func init() {

	if _, err := os.Stat(".env"); err != nil {
		log.Println("Error loading .env file")
	}
	log.Println("Loading .env file")

}
