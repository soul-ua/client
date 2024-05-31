package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type StatusBar struct {
	focused bool
	width   int
}

func NewStatusBar() StatusBar {
	return StatusBar{}
}

func (b *StatusBar) SetSize(width int) {
	b.width = width
}

func (b *StatusBar) Focus() {
	b.focused = true
}

func (b *StatusBar) Blur() {
	b.focused = false
}

func (b StatusBar) IsFocused() bool {
	return b.focused
}

func (b StatusBar) Init() tea.Cmd {
	return nil
}

func (b StatusBar) View() string {
	style := lipgloss.NewStyle().
		Width(b.width).
		Height(1)
	if b.focused {
		style = style.Background(lipgloss.Color("235"))
	} else {
		style = style.Background(lipgloss.Color("234"))
	}

	return style.Render("statusbar")
}

func (b StatusBar) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return b, nil
}
