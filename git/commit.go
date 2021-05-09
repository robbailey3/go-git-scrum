package git

import (
	"strconv"
	"strings"
	"time"
)

type Commit struct {
	Content string
	Date    time.Time
}

func ParseCommit(commitString string) Commit {
	tokens := strings.Split(commitString, " ")

	timestamp := 0

	if len(tokens) > 1 {
		timestamp, _ := findTimestampAndIndex(tokens)
	}

	return Commit{
		Date: time.Unix(timestamp, 0),
	}
}

func findTimestampAndIndex(tokens []string) (int64, int) {
	for i := 4; i < 7; i += 1 {
		timestamp, err := strconv.ParseInt(tokens[i], 10, 32)
		if err != nil {
			continue
		}

		return timestamp, i
	}
	return 0, 0
}
