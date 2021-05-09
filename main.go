package main

import (
	"log"
	"os"
	"time"

	"github.com/robbailey3/go-git-scrum/git"
)

func main() {
	wd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	repos := git.GetRepositories(wd)

	for _, repo := range repos {
		for _, branch := range repo.Branches {
			commits := branch.GetCommitsAfterDate(time.Now().Add(-3 * 24 * time.Hour))
			if len(commits) > 0 {
				for _, commit := range commits {
					commit.Print()
				}
			}

		}
	}
}
