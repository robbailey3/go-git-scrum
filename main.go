package main

import (
	"fmt"
	"log"
	"os"

	"github.com/robbailey3/go-git-scrum/git"
)

func main() {
	wd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	repos := git.GetRepositories(wd)

	for _, repo := range repos {
		fmt.Println(repo.Name)
		fmt.Println(repo.Path)
		for _, branch := range repo.GetBranches() {
			for _, commit := range branch.Commits {
				fmt.Println(commit.Date)
				fmt.Println("")
			}
		}
	}
}
