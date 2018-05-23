package main

import (
	"github.com/oreuta/elementary-service/src/services"
	"os"
	"log"
)

func main() {

	logFile, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	log.Println("Recording of the log file has started...")

	services.Router()
}



