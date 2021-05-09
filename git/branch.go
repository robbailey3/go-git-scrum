package git

import (
	"github.com/robbailey3/go-git-scrum/file"
	"strings"
	"time"
)

type Branch struct {
	RepoName string
	Name     string
	Path     string
	Commits  []*Commit
}

func NewBranch(name, path, repoName string) *Branch {
	branch := Branch{
		RepoName: repoName,
		Name:     name,
		Path:     path,
	}

	branch.getCommits()

	return &branch
}

func (b *Branch) getCommits() {
	str := string(file.Read(b.Path))

	commitStrings := strings.Split(str, "\n")

	var commits []*Commit

	for _, commitString := range commitStrings {
		commit := ParseCommit(commitString, b.Name, b.RepoName)
		if commit != nil {
			commits = append(commits, commit)
		}
	}

	b.Commits = commits
}

func (b *Branch) GetCommitsAfterDate(maxAge time.Time) []*Commit {
	return maxAgeFilter(b.Commits, maxAge)
}

func maxAgeFilter(commits []*Commit, maxAge time.Time) []*Commit {
	var result []*Commit
	for _, commit := range commits {
		if commit.Date.After(maxAge) {
			result = append(result, commit)
		}
	}
	return result
}
