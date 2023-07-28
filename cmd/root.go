package cmd

import (
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "go-hexa",
	Short: "hexagonal cli generator",
	Long: "go-hexa cli generates a hexagonal template written in go and fiber for http server and sqlboiler for ORM",
}

func Execute() error {
	return rootCommand.Execute()
}
