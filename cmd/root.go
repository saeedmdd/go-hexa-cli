package cmd

import (
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "go-hexa",
	Short: "hexagonal cli generetor",
}

func Execute() error {
	return rootCommand.Execute()
}
