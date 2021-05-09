package git

import (
	"fmt"
	"github.com/robbailey3/go-git-scrum/file"
	"log"
	"os"
	"time"
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

func PrintLatestCommits() {
	wd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	repos := GetRepositories(wd)

	for _, repo := range repos {
		for _, branch := range repo.Branches {
			fmt.Println(branch.Name)
			commits := branch.GetCommitsAfterDate(time.Now().Add(-7 * 24 * time.Hour))
			if len(commits) > 0 {
				for _, commit := range commits {
					commit.Print()
				}
			}

		}
	}
}
