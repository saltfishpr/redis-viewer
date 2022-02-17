// @file: update.go
// @date: 2022/02/08

package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// handleMouse handles all mouse interaction.
func (m *model) handleMouse(msg tea.MouseMsg) {
	switch msg.Type {
	case tea.MouseWheelUp:
		m.viewport.LineUp(mouseScrollSpeed)
		m.viewport.SetContent(m.detailView())
	case tea.MouseWheelDown:
		m.viewport.LineDown(mouseScrollSpeed)
		m.viewport.SetContent(m.detailView())
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
			case key.Matches(msg, m.keyMap.scan):
				m.ready = false
				cmds = append(cmds, m.scanCmd(), m.countCmd())
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

			m.ready = false
			cmds = append(cmds, m.scanCmd(), m.countCmd())
		default:
			m.textinput, cmd = m.textinput.Update(msg)
			cmds = append(cmds, cmd)
		}
	}

	return tea.Batch(cmds...)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case errMsg:
		m.statusMessage = msg.err.Error()
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		statusBarHeight := lipgloss.Height(m.statusView())
		height := m.height - statusBarHeight

		listViewWidth := int(listProportion * float64(m.width))
		listWidth := listViewWidth - listViewStyle.GetHorizontalFrameSize()
		m.list.SetSize(listWidth, height)

		detailViewWidth := m.width - listViewWidth
		m.viewport = viewport.New(detailViewWidth, height)
		m.viewport.SetContent(m.detailView())
	case scanMsg:
		m.list.SetItems(msg.items)
	case countMsg:
		if msg.count > maxScanCount {
			m.statusMessage = fmt.Sprintf("%d+ keys found", maxScanCount)
		} else {
			m.statusMessage = fmt.Sprintf("%d keys found", msg.count)
		}
		m.ready = true
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

		m.viewport, cmd = m.viewport.Update(msg)
		cmds = append(cmds, cmd)

		m.spinner, cmd = m.spinner.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
