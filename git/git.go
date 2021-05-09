package git

import (
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
