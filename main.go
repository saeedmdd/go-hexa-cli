package main

import (
	"log"

	"github.com/saeedmdd/go-hexa-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Panicf("error while executing app: %s", err)
	}
}
