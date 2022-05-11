package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

    log.SetOutput(file)

	fmt.Println("LifeSaver says Hi!")
	log.Println("Creating log statement.")

	initRouter()
}