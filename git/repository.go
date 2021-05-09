package git

import (
	"github.com/robbailey3/go-git-scrum/file"
)

type Repository struct {
	Name     string
	Path     string
	Branches []*Branch
}

func NewRepository(name string, path string) *Repository {
	repository := Repository{
		Name: name,
		Path: path,
	}

	repository.GetBranches()

	return &repository
}

func (r *Repository) GetBranches() {
	r.Branches = getBranches(r.Path+"\\.git\\logs\\refs\\heads", r.Name)
}

func getBranches(path string, repoName string) []*Branch {
	dirs := file.ReadDir(path)

	var result []*Branch

	for _, dir := range dirs {
		path := path + "\\" + dir.Name()
		if !dir.IsDir() {
			result = append(result, NewBranch(dir.Name(), path, repoName))
		} else {
			result = append(result, getBranches(path, repoName)...)
		}
	}
	return result
}
