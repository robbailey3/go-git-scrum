package ui

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	headerHeight = 3
	footerHeight = 3
)

type BaseUiModel struct {
	viewport viewport.Model
	ready    bool
	content  string
}

func InitBaseUiModel() *BaseUiModel {
	return &BaseUiModel{
		content: "Hello world",
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
			m.viewport.SetContent(m.content)
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMargins
		}
	}
	return m, nil
}

func (m *BaseUiModel) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}
	return m.viewport.View()
}
