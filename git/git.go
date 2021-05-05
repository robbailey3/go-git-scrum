package git

import (
	"fmt"
	"github.com/robbailey3/go-git-scrum/file"
)

func GetCommits(path string) {

}

func GetGitRepos(path string) []string {
	dirs := file.ReadDir(path)
	var result []string
	for _, dir := range dirs {
		fmt.Println(dir.Name())
		if file.Exists(path + "\\" + dir.Name() + "\\.git/logs/refs/heads") {
			result = append(result, path+"\\"+dir.Name())
		}
	}
	return result
}

func GetRepoCommits(path string) []string {
	branches := file.ReadDir(path)

	var result []string

	for _, branch := range branches {
		fmt.Println(branch.Name())
		if branch.IsDir() {
			result = append(result, GetGitRepos(path+"\\"+branch.Name())...)
		} else {
			result = append(result, string(file.Read(path+"\\"+branch.Name())))
		}
	}

	return result
}
