package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Setting up workspace",
	Long:  "Generate workspace directory and config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initialize is done!")
	},
}
