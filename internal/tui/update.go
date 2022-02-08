// @description:
// @file: update.go
// @date: 2022/02/08

package tui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type state int

const (
	defaultState state = iota
	insertState
)

// handleMouse handles all mouse interaction.
func (m *model) handleMouse(msg tea.MouseMsg) {
	switch msg.Type {
	case tea.MouseWheelUp:
		m.cursor = (m.cursor + len(m.choices) - 1) % len(m.choices)
	case tea.MouseWheelDown:
		m.cursor = (m.cursor + len(m.choices) + 1) % len(m.choices)
	}
}

// handleKeys handles all keypresses.
func (m *model) handleKeys(msg tea.KeyMsg) tea.Cmd {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch {
	case key.Matches(msg, m.keyMap.Quit):
		return tea.Quit
	case key.Matches(msg, m.keyMap.Up):
		if m.state == defaultState {
			m.cursor = (m.cursor + len(m.choices) - 1) % len(m.choices)
		}
	case key.Matches(msg, m.keyMap.Down):
		if m.state == defaultState {
			m.cursor = (m.cursor + len(m.choices) + 1) % len(m.choices)
		}
	case key.Matches(msg, m.keyMap.Space):
		_, ok := m.selected[m.cursor]
		if ok {
			delete(m.selected, m.cursor)
		} else {
			m.selected[m.cursor] = struct{}{}
		}
	case key.Matches(msg, m.keyMap.Insert):
		if m.state == defaultState {
			m.state = insertState
			m.textinput.Placeholder = "Enter new choice"
			m.textinput.Focus()
			return textinput.Blink
		}
	case key.Matches(msg, m.keyMap.Enter):
		switch m.state {
		case insertState:
			m.choices = append(m.choices, m.textinput.Value())
			m.textinput.Blur()
			m.textinput.Reset()
			m.state = defaultState
		default:

		}
	case key.Matches(msg, m.keyMap.Escape):
		m.state = defaultState
	}
	m.textinput, cmd = m.textinput.Update(msg)
	cmds = append(cmds, cmd)

	return tea.Batch(cmds...)
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		return m, nil
	case tea.MouseMsg:
		m.handleMouse(msg)
	case tea.KeyMsg:
		cmd = m.handleKeys(msg)
		cmds = append(cmds, cmd)

		return m, tea.Batch(cmds...)
	}

	m.textinput, cmd = m.textinput.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
