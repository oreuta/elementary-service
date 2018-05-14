package main

import (
	"log"
	"github.com/elementary-service/src/services"
)

func main() {
	log.Println("Elementary-service started...")
	services.Router()
}
