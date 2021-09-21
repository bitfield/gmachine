package main

import (
	"gmachine"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Usage: run [gbin file]\n")
	}
	gmachine.RunCLI(os.Args[1])
}
