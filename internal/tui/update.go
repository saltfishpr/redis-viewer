// @description:
// @file: update.go
// @date: 2022/02/08

package tui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// handleMouse handles all mouse interaction.
func (m *model) handleMouse(msg tea.MouseMsg) {
	switch msg.Type {
	case tea.MouseWheelUp:
		m.list.CursorUp()
	case tea.MouseWheelDown:
		m.list.CursorDown()
	}
}

// handleKeys handles all keypresses.
func (m *model) handleKeys(msg tea.KeyMsg) tea.Cmd {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch m.state {
	case defaultState: // default state, focus on list
		switch msg.Type {
		case tea.KeyRunes:
			switch {
			case key.Matches(msg, m.keyMap.search):
				m.state = searchState
				m.textinput.Focus()
				return textinput.Blink
			}
		case tea.KeyCtrlC:
			cmd = tea.Quit
			cmds = append(cmds, cmd)
		default:
			m.list, cmd = m.list.Update(msg)
			cmds = append(cmds, cmd)
		}
	case searchState: // search state, focus on textinput
		switch msg.Type {
		case tea.KeyRunes:
			m.textinput, cmd = m.textinput.Update(msg)
			cmds = append(cmds, cmd)
		case tea.KeyEscape:
			m.textinput.Blur()
			m.textinput.Reset()
			m.state = defaultState
		case tea.KeyEnter:
			m.searchValue = m.textinput.Value()
			m.textinput.Blur()
			m.textinput.Reset()
			m.state = defaultState
		default:
			m.textinput, cmd = m.textinput.Update(msg)
			cmds = append(cmds, cmd)
		}
	}

	return tea.Batch(cmds...)
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		topGap, rightGap, bottomGap, leftGap := appStyle.GetPadding()
		m.list.SetSize(msg.Width-leftGap-rightGap, msg.Height-topGap-bottomGap-1)
	case tea.MouseMsg:
		m.handleMouse(msg)
	case tea.KeyMsg:
		cmd = m.handleKeys(msg)
		cmds = append(cmds, cmd)
	default:
		m.list, cmd = m.list.Update(msg)
		cmds = append(cmds, cmd)

		m.textinput, cmd = m.textinput.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
