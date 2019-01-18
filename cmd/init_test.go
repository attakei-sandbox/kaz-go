package cmd

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExecute(t *testing.T) {
	assert := require.New(t)
	dir, err := ioutil.TempDir("", "kaz-cmd-init")
	if err != nil {
		assert.Fail("Failure create temp dir")
	}
	defer os.RemoveAll(dir)

	Execute()
}

func TestCreateWorkDirs(t *testing.T) {
	assert := require.New(t)
	dir, err := ioutil.TempDir("", "kaz-cmd-init")
	if err != nil {
		assert.Fail("Failure create temp dir")
	}
	defer os.RemoveAll(dir)

	createWorkDirs(dir)
	assert.DirExists(path.Join(dir, ".kaz"))
	assert.DirExists(path.Join(dir, ".kaz", "bin"))
	assert.DirExists(path.Join(dir, ".kaz", "log"))
	assert.DirExists(path.Join(dir, ".kaz", "repos"))
}

func TestOutputNextMessage(t *testing.T) {
	assert := require.New(t)
	w := new(bytes.Buffer)
	outputNextMessage(w, "/home/attakei")
	assert.Contains(w.String(), "/home/attakei/.kaz/bin")
}

func TestCreateDefaultConfig(t *testing.T) {
	assert := require.New(t)
	dir, err := ioutil.TempDir("", "kaz-cmd-init")
	if err != nil {
		assert.Fail("Failure create temp dir")
	}
	defer os.RemoveAll(dir)

	confPath := path.Join(dir, "kaz.cnf")
	createDefaultConfig(confPath)
	assert.FileExists(confPath)
}
