package main

import (
	"io/ioutil"
	"os"
)

func Files(dir string) ([]os.FileInfo, error) {
	fileInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var files []os.FileInfo
	for _, f := range fileInfo {
		if !f.IsDir() {
			files = append(files, f)
		}
	}

	return files, nil
}
