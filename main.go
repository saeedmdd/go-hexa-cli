package main

import (
	"fmt"

	"github.com/saeedmdd/go-hexa-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err.Error())
	}
}
