package ui

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/robbailey3/go-git-scrum/git"
)

const (
	headerHeight = 3
	footerHeight = 3
)

type BaseUiModel struct {
	viewport     viewport.Model
	ready        bool
	content      string
	Repositories []*git.Repository
}

func InitBaseUiModel(repos []*git.Repository) *BaseUiModel {
	return &BaseUiModel{
		Repositories: repos,
	}
}

func (m *BaseUiModel) Init() tea.Cmd {
	return nil
}

func (m *BaseUiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// var (
	// 	cmd  tea.Cmd
	// 	cmds []tea.Cmd
	// )
	m.content = ""
	for _, repo := range m.Repositories {
		m.content += repo.Name + "\n"
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		verticalMargins := headerHeight + footerHeight

		if !m.ready {
			m.viewport = viewport.Model{Width: msg.Width, Height: msg.Height - verticalMargins}
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMargins
		}
	}
	m.viewport.SetContent(m.content)
	return m, nil
}

func (m *BaseUiModel) View() string {
	if !m.ready {
		return "Initializing..."
	}
	if len(m.Repositories) == 0 {
		return "No repositories found"
	}
	return m.viewport.View()
}
