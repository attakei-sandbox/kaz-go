package cmd

import (
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
