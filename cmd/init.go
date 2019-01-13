package cmd

import (
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var defaultConfig string = `
########################################
# KAZ cofigs                           #
########################################

# ------------------
# Base directory settings
# ------------------
bin_dir = "/opt/kaz/bin"
log_dir = "/var/log/kaz"
app_dir = "/var/opt/kaz"
`

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize environment to work",
	Long: `Initialize workspace.

KAZ need workspace directory to save appimage and history.
'init' command crete directory at first.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create default config(if not specifed)
		cfgFile = "/etc/kaz.cfg" // TODO: Fix after?
		if viper.ConfigFileUsed() == cfgFile {
			if _, e := os.Stat(cfgFile); os.IsNotExist(e) {
				fmt.Println("Default config file is not found.")
				fmt.Println("Create config ...")
				if cfg, err := os.Create(cfgFile); err != nil {
					fmt.Println(err)
					os.Exit(1)
				} else if tmpl, err := template.New("").Parse(defaultConfig); err != nil {
					fmt.Println(err)
					os.Exit(1)
				} else if err := tmpl.Execute(cfg, nil); err != nil {
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
