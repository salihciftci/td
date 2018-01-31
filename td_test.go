package main

import (
	"os"
	"testing"
)

func TestDir(t *testing.T) {
	path := os.Getenv("HOME") + "/.tddb"
	if _, err := os.Stat(os.Getenv(path)); os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			t.Error("Cannot write to HOME dir")
		}
		defer file.Close()
	}
}
