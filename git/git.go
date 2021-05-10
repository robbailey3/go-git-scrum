package git

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/robbailey3/go-git-scrum/file"
)

func GetRepositories(path string) []*Repository {
	dirs := file.ReadDir(path)
	var result []*Repository
	for _, dir := range dirs {
		if file.Exists(path + "\\" + dir.Name() + "\\.git/logs/refs/heads") {
			result = append(result, NewRepository(dir.Name(), path+"\\"+dir.Name()))
		}
	}
	return result
}

func PrintLatestCommits(numberOfDays int) {
	wd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	repos := GetRepositories(wd)

	for _, repo := range repos {
		for _, branch := range repo.Branches {
			commits := branch.GetCommitsAfterDate(time.Now().Add(time.Duration(-numberOfDays*24) * time.Hour))
			if len(commits) > 0 {
				fmt.Println(repo.Name)
				for _, commit := range commits {
					commit.Print()
				}
			}
		}
	}
}
