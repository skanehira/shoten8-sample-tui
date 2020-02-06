package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func mkfile(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

func mkdir(name string) error {
	return os.Mkdir(name, 0666)
}

func TestFiles(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		testdir, err := ioutil.TempDir("", "")
		if err != nil {
			t.Fatalf("cannot create testdir: %s", err)
		}
		defer os.RemoveAll(testdir)

		exceptedFiles := map[string]string{
			"a.go": "f",
			"b.md": "f",
			"tmp":  "d",
		}

		for f, typ := range exceptedFiles {
			// if file
			if typ == "f" {
				if err := mkfile(filepath.Join(testdir, f)); err != nil {
					t.Fatalf("cannot create test file: %s", err)
				}
				// if dir
			} else if typ == "d" {
				if err := mkdir(filepath.Join(testdir, f)); err != nil {
					t.Fatalf("cannot create test dir: %s", err)
				}
			}
		}

		files, err := files(testdir)
		if err != nil {
			t.Fatalf("cannot get files: %s", err)
		}

		for _, f := range files {
			if _, ok := exceptedFiles[f.Name()]; !ok {
				t.Fatalf("unexcepted file, want: a.go or b.md, got: %s", f.Name())
			}
		}
	})

	t.Run("failed", func(t *testing.T) {
		if _, err := files("xxx"); err == nil {
			t.Fatalf("unexcepted test: err is nil")
		}
	})
}
