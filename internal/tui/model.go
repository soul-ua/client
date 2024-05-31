package tui

import (
	"os"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
)

type model struct {
	chats       Box
	currentChat Box
	inbox       Box
	statusBar   StatusBar
}

func initialModel() model {
	m := model{
		chats:       NewBox(0, 0, nil),
		currentChat: NewBox(0, 0, nil),
		inbox:       NewBox(0, 0, nil),
		statusBar:   NewStatusBar(),
	}
	m.chats.Focus()
	return m
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
	)

	//m.textarea, tiCmd = m.textarea.Update(msg)
	//m.viewport, vpCmd = m.viewport.Update(msg)
	//
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEsc:
			if m.statusBar.IsFocused() {
				m.statusBar.Blur()
				m.chats.Focus()
			} else {
				m.chats.Blur()
				m.currentChat.Blur()
				m.inbox.Blur()
				m.statusBar.Focus()
			}
			return m, nil
		case tea.KeyShiftTab:
			if m.chats.IsFocused() {
				m.chats.Blur()
				m.inbox.Focus()
				return m, nil
			}
			if m.currentChat.IsFocused() {
				m.currentChat.Blur()
				m.chats.Focus()
				return m, nil
			}
			if m.inbox.IsFocused() {
				m.inbox.Blur()
				m.currentChat.Focus()
				return m, nil
			}
		case tea.KeyTab:
			if m.statusBar.IsFocused() {
				m.statusBar.Blur()
				m.chats.Focus()
			} else if m.chats.IsFocused() {
				m.chats.Blur()
				m.currentChat.Focus()
			} else if m.currentChat.IsFocused() {
				m.currentChat.Blur()
				m.inbox.Focus()
			} else if m.inbox.IsFocused() {
				m.inbox.Blur()
				m.chats.Focus()
			}
			return m, nil
		}

	// We handle errors just like any other message
	case errMsg:
		//m.err = msg
		return m, nil
	}

	return m, tea.Batch(tiCmd, vpCmd)
}

func (m model) View() string {
	physicalWidth, physicalHeight, _ := term.GetSize(os.Stdout.Fd())

	viewPortHeight := physicalHeight - 1
	m.chats.SetSize(30, viewPortHeight)
	m.inbox.SetSize(40, viewPortHeight)
	m.currentChat.SetSize(physicalWidth-30-40, viewPortHeight)
	m.statusBar.SetSize(physicalWidth)

	return lipgloss.JoinVertical(lipgloss.Top,
		lipgloss.JoinHorizontal(lipgloss.Top,
			m.chats.View(), m.currentChat.View(), m.inbox.View(),
		),
		m.statusBar.View(),
	)
}
