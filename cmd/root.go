package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/robbailey3/go-git-scrum/ui"
	"github.com/spf13/cobra"
)

var numberOfDays int

var rootCmd = &cobra.Command{
	Use:   "go-git-scrum",
	Short: "A command line interface to show latest commits",
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(ui.InitBaseUiModel(), tea.WithAltScreen(), tea.WithMouseCellMotion())
		if err := p.Start(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	rootCmd.PersistentFlags().IntVarP(&numberOfDays, "Number of days", "n", 3, "The number of days to show commits for")
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
