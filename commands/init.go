package commands

import (
	"fmt"
	"log"
	"os"
	"path"
	"text/template"

	"github.com/gobuffalo/packr"
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
		// Generate work directories
		makeWorkDirs("/") // TODO: Accept custom dir
		// Generate config(if not argumented)
		makeDefaultConf("/")
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

// Make default config file
func makeDefaultConf(basePath string) error {
	box := packr.NewBox("../assets")
	confPath := path.Join(basePath, "etc/kaz.conf")
	_ = os.Mkdir(path.Join(basePath, "etc"), 0755)
	tmplStr, err := box.FindString("kaz.conf.tpl")
	tmpl, err := template.New("kaz-conf").Parse(tmplStr)
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
