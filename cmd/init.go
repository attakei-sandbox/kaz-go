// Copyright © 2019 Kazuya Takei <attakei@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
)

type InitParam struct {
	BaseDir    string
	WorkDir    string
	ConfigPath string
}

func NewParam(baseDir string) InitParam {
	p := InitParam{}
	p.BaseDir = baseDir
	p.WorkDir = filepath.Join(p.BaseDir, ".kaz")
	p.ConfigPath = filepath.Join(p.WorkDir, "kaz.cnf")
	return p
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize kaz working space",
	Long:  `Initialize working space and instructs next action for user.`,
	Run: func(cmd *cobra.Command, args []string) {
		param := NewParam(os.Getenv("HOME"))
		if err := createWorkDirs(param.BaseDir); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		createDefaultConfig(filepath.Join(param.BaseDir, "kaz.cfg"), param)
		outputNextMessage(os.Stdout, param)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

// TODO: More message
func outputNextMessage(writer io.Writer, param InitParam) error {
	const msg = `Congratulations!!
You can manage applications by kaz.

Application is installed at {{.WorkDir}}/bin
Set PATH into it

export PATH={{.workDir}}/bin:$PATH
`
	if tmpl, err := template.New("init-output").Parse(msg); err != nil {
		return err
	} else if err := tmpl.Execute(writer, param); err != nil {
		return err
	}
	return nil
}

func createWorkDirs(baseDir string) error {
	targets := []string{
		".kaz",
		".kaz/log",
		".kaz/bin",
		".kaz/repos",
		".kaz/data",
	}
	for _, target := range targets {
		if err := os.Mkdir(filepath.Join(baseDir, target), 0700); err != nil {
			return err
		}
	}
	return nil
}

func createDefaultConfig(target string, param InitParam) error {
	// TODO: file templating after
	const contentTmpl = `# ------
# kaz config
# -----
work_dir = {{.BaseDir}}
`
	if tmpl, err := template.New("default-config").Parse(contentTmpl); err != nil {
		return err
	} else if f, err := os.Create(target); err != nil {
		return err
	} else if err := tmpl.Execute(f, param); err != nil {
		return err
	}
	return nil
}
