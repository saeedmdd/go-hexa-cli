package cmd

import (
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/saeedmdd/go-hexa-cli/internal/hexagonal"
	"github.com/saeedmdd/go-hexa-cli/internal/hexagonal/modifier"
	"github.com/saeedmdd/go-hexa-cli/utils"
	"github.com/spf13/cobra"
)

var newCommand = &cobra.Command{
	Use:   "new",
	Short: "generates a new hexagonal template in the given directory",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Println(cmd.Help())
			return
		}

		exists, err := utils.DirectoryExists(args[0])
		if err != nil {
			cmd.Print("error while executing command\n")
		}
		if exists {
			cmd.PrintErrf("directory %s already exists\n", args[0])
			return
		}

		rootFigure := figure.NewColorFigure("Go hexa", "", "cyan", true)
		rootFigure.Print()
		monolithValue, _ := cmd.Flags().GetBool("monolith")
		branchName := "main"
		if monolithValue {
			branchName = "monolith"
		}
		err = os.Mkdir(args[0], 0755)
		if err != nil {
			cmd.PrintErrf("error while making directory: %v", err)
		}
		hexa := hexagonal.NewHexagonal("https://github.com/professionsforall/hexagonal-template.git", args[0], branchName)

		err = hexa.Generate(cmd.Context())
		if err != nil {
			panic(err)
		}
		dotGitFile := fmt.Sprintf("%s/.git", args[0])
		err = os.RemoveAll(dotGitFile)
		if err != nil {
			panic(err)
		}
		modValue, _ := cmd.Flags().GetString("mod")
		if modValue != "" {
			modifier.ModifyFiles("./"+args[0], "github.com/professionsforall/hexagonal-template", modValue)
		}

	},
	Example: "go-hexa-cli new <directory> <flags>",
}

func init() {
	rootCommand.AddCommand(newCommand)
	newCommand.PersistentFlags().StringP(
		"mod",
		"m",
		"",
		"set the mod path for your hexagonal project")
	newCommand.Flags().Bool("monolith", false, "to use the architecture as monolith project")
}
