package commands

import (
	"fmt"
	"log"
	"os"
	"path"
	"text/template"

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
		makeDefaultConf("/")
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

// Make default config file
func makeDefaultConf(basePath string) error {
	confPath := path.Join(basePath, "etc/kaz.conf")
	_ = os.Mkdir(path.Join(basePath, "etc"), 0755)
	tmpl, err := template.New("kaz-default-config").Parse("# kaz config")
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(confPath)
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(f, nil)
	f.Close()
	return nil
}
