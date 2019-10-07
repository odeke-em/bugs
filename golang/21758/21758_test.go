package filepath_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	dir, err := ioutil.TempDir("/tmp/", "fptest")
	if err != nil {
		t.Fatalf("TempDir failed: %v", err)
	}

	crdir := filepath.Join(dir, "cant-read")
	if err := os.Mkdir(crdir, 0300); err != nil {
		t.Fatalf("Mkdir(%s) failed: %v", err)
	}

	var got []string
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err == nil {
			got = append(got, path)
		} else {
t.Logf("err=%v", err)
}
		return nil
	})
	if err != nil {
		t.Errorf("Walk(%s) failed: %v", dir, dir)
	}

	want := []string{dir, crdir}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Walk was called with paths %s; want %s", got, want)
	}
}
