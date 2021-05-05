package main

import (
	"fmt"
	"log"
	"os"

	"github.com/robbailey3/go-git-scrum/file"
)

func main() {
	wd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	dirs := file.ReadDir(wd)

	for _, dir := range dirs {
		fmt.Println(dir.Name())
	}
}
