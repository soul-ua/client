package tui

// A simple program demonstrating the text area component from the Bubbles
// component library.

import (
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

func Run() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type (
	errMsg error
)
