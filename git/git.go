package git

import (
	"fmt"
	"strconv"
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

func GetRepositoriesWithCommits(n int, path string) []*Repository {

	repos := GetRepositories(path)
	return filterRepos(repos, func(repo *Repository) bool {
		hasNewCommit := false
		for _, branch := range repo.Branches {
			if len(branch.Commits) < 1 {
				continue
			}
			for _, commit := range branch.Commits {
				if commit.TimeSince < time.Duration(n*int(time.Hour)*24) {
					hasNewCommit = true
				}
			}
		}
		return hasNewCommit
	})
}

func filterRepos(repos []*Repository, filterFunc func(repo *Repository) bool) []*Repository {
	var result []*Repository
	for _, repo := range repos {
		if filterFunc(repo) {
			result = append(result, repo)
		}
		fmt.Println(repo.Name + " " + strconv.FormatBool(filterFunc(repo)))
	}
	return result
}
