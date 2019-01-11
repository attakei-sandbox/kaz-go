package commands

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestMakeWorkDirs_ok(t *testing.T) {
	dir, err := ioutil.TempDir("", "kaz.testing.")
	defer os.RemoveAll(dir)
	if err != nil {
		t.Fatal(err)
	}
	if err = makeWorkDirs(dir); err != nil {
		t.Fatal(err)
	}
	if _, err = os.Stat(dir + "/usr/local/bin"); err != nil {
		t.Fatal(err)
	}
	if _, err = os.Stat(dir + "/var/opt/kaz"); err != nil {
		t.Fatal(err)
	}
	if _, err = os.Stat(dir + "/var/log/kaz"); err != nil {
		t.Fatal(err)
	}
}
