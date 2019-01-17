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
	"path"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize kaz working space",
	Long:  `Initialize working space and instructs next action for user.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

// TODO: More message
func outputNextMessage(writer io.Writer, baseDir string) {
	fmt.Fprintf(writer, `Congratulations!!
You can manage applications by kaz.

Application is installed at %s/.kaz/bin
Set PATH into it

export PATH=%s/.kaz/bin:$PATH
`, baseDir, baseDir)
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
		if err := os.Mkdir(path.Join(baseDir, target), 0700); err != nil {
			return err
		}
	}
	return nil
}
