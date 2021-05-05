package file

import (
	"errors"
	"log"
	"os"
)

func Exists(path string) bool {
	_, err := os.Stat(path)

	return errors.Is(err, os.ErrNotExist)
}

func Read(path string) []byte {
	file, err := os.ReadFile(path)

	if err != nil {
		handleError(err)
	}

	return file
}

func ReadDir(path string) []os.DirEntry {
	dirs, err := os.ReadDir(path)

	if err != nil {
		handleError(err)
	}

	return dirs
}

func handleError(err error) {
	log.Fatal(err)
}
