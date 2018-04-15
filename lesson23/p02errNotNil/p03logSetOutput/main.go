package main

import (
	"fmt"
	"log"
	"os"
)

func init() {

	logFile, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		if err != nil {
			fmt.Println(err)
		}
	}

	log.SetOutput(logFile)
}

func main() {

	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("err happened", err)
	}
}
