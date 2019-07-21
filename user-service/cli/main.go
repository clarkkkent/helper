package main

import (
	"flag"
	"log"
)

const (
	address = "localhost:50051"
)

func main() {
	namePtr := flag.String("username", "", "name for user")
	emailPtr := flag.String("email", "", "email for user")
	passwordPtr := flag.String("password", "", "password for user")
	flag.Parse()

	log.Println(*namePtr, *emailPtr, *passwordPtr)
}

