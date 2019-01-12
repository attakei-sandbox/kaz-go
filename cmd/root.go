package cmd

import (
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string = "/etc/kaz.cfg"

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

var rootCmd = &cobra.Command{
	Use:   "kaz",
	Short: "Simple AppImage manager",
	Long:  "Simple AppImage manager",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// If default config file is not exist, generate it (need permission)
	if _, e := os.Stat(cfgFile); os.IsNotExist(e) {
		fmt.Println("Default config file is not found.")
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
	}
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetConfigType("toml")
	viper.SetConfigFile(cfgFile)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
	} else {
		fmt.Println("Using config file: ", viper.ConfigFileUsed())
	}
}
