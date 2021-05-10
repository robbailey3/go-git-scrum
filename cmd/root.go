package cmd

import (
	"fmt"
	"os"

	"github.com/robbailey3/go-git-scrum/git"
	"github.com/spf13/cobra"
)

var numberOfDays int

var rootCmd = &cobra.Command{
	Use:   "go-git-scrum",
	Short: "A command line interface to show latest commits",
	Run: func(cmd *cobra.Command, args []string) {
		git.PrintLatestCommits(numberOfDays)
	},
}

func Execute() {
	rootCmd.PersistentFlags().IntVarP(&numberOfDays, "Number of days", "n", 3, "The number of days to show commits for")
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
