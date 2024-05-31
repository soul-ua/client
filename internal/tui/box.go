package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Box struct {
	width    int
	height   int
	focused  bool
	children tea.Model
}

func NewBox(width, height int, children tea.Model) Box {
	return Box{
		width:    width,
		height:   height,
		children: children,
	}
}

func (b *Box) SetSize(width, height int) {
	b.width = width
	b.height = height
}

func (b *Box) Focus() {
	b.focused = true
}

func (b *Box) Blur() {
	b.focused = false
}

func (b Box) IsFocused() bool {
	return b.focused
}

func (b Box) Init() tea.Cmd {
	return nil
}

func (b Box) View() string {
	style := lipgloss.NewStyle().
		Width(b.width - 2).
		Height(b.height - 2).
		BorderStyle(lipgloss.RoundedBorder())
	if b.focused {
		style = style.BorderForeground(lipgloss.Color("63"))
	} else {
		style = style.BorderForeground(lipgloss.Color("60"))
	}

	return style.Render("box")

}

func (b Box) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return b, nil
}
