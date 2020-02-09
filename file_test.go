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

func TestFiles(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		testdir, err := ioutil.TempDir("", "")
		if err != nil {
			t.Fatalf("cannot create testdir: %s", err)
		}
		defer os.RemoveAll(testdir)

		exceptedFiles := map[string]string{
			"a.go": "f",
			"tmp":  "d",
		}

		for f, typ := range exceptedFiles {
			tmp := filepath.Join(testdir, f)
			// if file
			if typ == "f" {
				err := mkfile(tmp)
				if err != nil {
					t.Fatalf("create error: %s", err)
				}
				// if dir
			} else if typ == "d" {
				err := os.Mkdir(tmp, 0666)
				if err != nil {
					t.Fatalf("create error: %s", err)
				}
			}
		}

		files, err := Files(testdir)
		if err != nil {
			t.Fatalf("cannot get files: %s", err)
		}

		fileName := files[0].Name()
		if fileName != "a.go" {
			t.Fatalf("want: a.go, got: %s", fileName)
		}
	})

	t.Run("failed", func(t *testing.T) {
		if _, err := Files("xxx"); err == nil {
			t.Fatalf("failed test: err is nil")
		}
	})
}
