package git

import (
	"github.com/fatih/color"
	"strconv"
	"strings"
	"time"
)

type Commit struct {
	indexOfDate int
	rawString   string
	tokens      []string
	RepoName    string
	Hash        string
	Name        string
	Message     string
	Date        time.Time
	TimeSince   time.Duration
	BranchName  string
}

func ParseCommit(commitString, branchName, repoName string) *Commit {
	tokens := strings.Split(commitString, " ")

	if len(tokens) > 1 {
		commit := Commit{
			rawString:  commitString,
			tokens:     tokens,
			BranchName: branchName,
			RepoName:   repoName,
		}
		commit.parseCommitDate()
		commit.calculateTimeSince()
		commit.parseCommitMessage()
		commit.parseName()
		return &commit
	}
	return nil
}

func (c *Commit) parseCommitDate() {
	for i := 4; i < 7; i++ {
		unixTimestamp, err := strconv.ParseInt(c.tokens[i], 10, 64)
		if err != nil {
			continue
		}
		c.indexOfDate = i
		c.Date = time.Unix(unixTimestamp, 0)
	}
}

func (c *Commit) parseCommitMessage() {
	message := ""
	for i := c.indexOfDate + 2; i < len(c.tokens); i++ {
		message += c.tokens[i] + " "
	}
	c.Message = strings.Trim(message, " ")
}

func (c *Commit) parseName() {
	name := ""
	for i := 2; i < c.indexOfDate-1; i++ {
		name += c.tokens[i] + " "
	}
	c.Name = strings.Trim(name, " ")
}

func (c *Commit) parseHash() {
	c.Hash = c.tokens[1]
}

func (c *Commit) calculateTimeSince() {
	c.TimeSince = time.Since(c.Date)
}

func (c *Commit) Print() {
	red := color.New(color.FgRed).PrintfFunc()
	blue := color.New(color.FgCyan).PrintfFunc()
	white := color.New(color.FgWhite).PrintfFunc()
	whiteUnderline := color.New(color.FgWhite).Add(color.Underline).PrintfFunc()

	whiteUnderline("%s --- %s \n", c.RepoName, c.BranchName)
	white("%s ", c.Message)
	blue("Time Since: %s ", c.TimeSince.String())
	red("Name: %s \n\n", c.Name)
}
