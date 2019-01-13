package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetConfigType("toml")
	viper.SetConfigFile(defaultConfigPath)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
	} else {
		fmt.Println("Using config file: ", viper.ConfigFileUsed())
	}
}
