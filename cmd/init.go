package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize environment to work",
	Long: `Initialize workspace.

KAZ need workspace directory to save appimage and history.
'init' command crete directory at first.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Create directories...")
		targets := []string{
			viper.GetString("bin_dir"),
			viper.GetString("log_dir"),
			viper.GetString("app_dir"),
		}
		for _, target := range targets {
			fmt.Println("- " + target)
			os.MkdirAll(target, 0755)
		}
		fmt.Println("Initialized!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
