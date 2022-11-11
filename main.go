package main

import (
	"log"

	"github.com/ant1k9/api-crawler/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
