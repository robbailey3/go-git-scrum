package git

import (
	"github.com/robbailey3/go-git-scrum/file"
)

type Repository struct {
	Name string
	Path string
}

func NewRepository(name string, path string) Repository {
	return Repository{
		Name: name,
		Path: path,
	}
}

func (r *Repository) GetBranches() []Branch {
	return getBranches(r.Path + "\\.git\\logs\\refs\\heads")
}

func getBranches(path string) []Branch {
	dirs := file.ReadDir(path)

	var result []Branch

	for _, dir := range dirs {
		path := path + "\\" + dir.Name()
		if !dir.IsDir() {
			result = append(result, NewBranch(dir.Name(), path))
		} else {
			result = append(result, getBranches(path)...)
		}
	}
	return result
}
