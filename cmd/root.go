package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kaz",
	Short: "Simple AppImage manager",
	Long: `kaz is simple AppImage manager. (pre-alpha version)

Please see https://github.com/attakei/kaz
`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
