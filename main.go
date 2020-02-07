package main

import (
	"fmt"
	"os"
)

func main() {
	if err := NewGUI().Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
