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
		// Create default config(if not specifed)
		if viper.ConfigFileUsed() == defaultConfigPath {
			if _, e := os.Stat(defaultConfigPath); os.IsNotExist(e) {
				fmt.Println("Default config file is not found.")
				fmt.Println("Create config ...")
				if err := createConfigFile(defaultConfigPath); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				fmt.Println("Created")
				viper.ReadInConfig()
			}
		}
		// Create working directories
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
		// Done
		fmt.Println("Initialized!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
