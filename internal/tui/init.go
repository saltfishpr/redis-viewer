// @file: init.go
// @date: 2022/02/08

package tui

import tea "github.com/charmbracelet/bubbletea"

func (m model) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, m.scanCmd(), m.countCmd())
}
