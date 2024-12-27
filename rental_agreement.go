package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		printUsage()
		os.Exit(1)
	}

	err := Checkout(args[0], args[1], args[2], args[3])
	if err != nil {
		printUsage()
		log.Fatal(err)
	}
}

func printUsage() {
	fmt.Println("Usage: rental_agreement <tool code> <number of rental days> <discount percent> <check out date>")
}
