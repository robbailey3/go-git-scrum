package git

import (
	"github.com/robbailey3/go-git-scrum/file"
	"strings"
)

type Branch struct {
	Name    string
	Path    string
	Commits []Commit
}

func NewBranch(name, path string) Branch {
	return Branch{
		Name:    name,
		Path:    path,
		Commits: getCommits(path),
	}
}

func getCommits(path string) []Commit {
	str := string(file.Read(path))

	commitStrings := strings.Split(str, "\n")

	var result []Commit

	for _, commitString := range commitStrings {
		result = append(result, ParseCommit(commitString))
	}

	return result
}
