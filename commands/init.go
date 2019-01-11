package commands

import (
	"fmt"
	"log"
	"os"
	"path"

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
		// Generate config(if not argumented)
		// Generate work directories
		makeWorkDirs("/") // TODO: Accept custom dir
		// Make base files
		// Display finish message
		fmt.Println("Initialize is done!")
	},
}

// Make necessary folders as working directory.
func makeWorkDirs(basePath string) error {
	targets := []string{
		"usr/local/bin",
		"var/opt/kaz",
		"var/log/kaz",
	}
	for _, p := range targets {
		targetPath := path.Join(basePath, p)
		if err := os.MkdirAll(targetPath, 0755); err != nil {
			log.Fatal(err)
		}
	}
	return nil
}
