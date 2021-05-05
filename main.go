package main

import (
	"fmt"
	"github.com/robbailey3/go-git-scrum/git"
	"log"
	"os"
)

func main() {
	wd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	repos := git.GetGitRepos(wd)

	var commits []string

	for _, repo := range repos {
		commits = append(commits, git.GetRepoCommits(repo+"\\.git\\logs\\refs\\heads")...)
	}
	for _, commit := range commits {
		fmt.Println(commit)
	}
}
