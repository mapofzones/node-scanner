package main

import (
	"log"
	"os"
)

func main() {
	rpc := os.Getenv("rpc")
	if len(rpc) == 0 {
		log.Fatal("The \"rpc\" env var cannot be empty!")
	}
	log.Println(rpc)
}
