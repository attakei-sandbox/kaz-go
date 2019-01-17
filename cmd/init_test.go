package cmd

import (
	"io/ioutil"
	"os"
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
