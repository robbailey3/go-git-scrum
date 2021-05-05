package file

import (
	"errors"
	"log"
	"os"
)

func Exists(path string) bool {
	_, err := os.Stat(path)

	return !errors.Is(err, os.ErrNotExist)
}

func IsDir(path string) bool {
	info, err := os.Stat(path)

	if err != nil {
		handleError(err)
	}

	return info.IsDir()
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
