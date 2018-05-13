package main

import (
	"log"
	"services"
)

func main() {
	log.Println("Elementary-service started...")
	services.Router()
}
