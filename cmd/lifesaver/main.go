package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)

	fmt.Println("LifeSaver says Hi!")

	initRouter()
}