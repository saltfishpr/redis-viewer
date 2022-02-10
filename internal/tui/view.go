// @description:
// @file: view.go
// @date: 2022/02/08

package tui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	appStyle = lipgloss.NewStyle().Padding(1, 2)
)

func (m *model) View() string {
	var state string
	var value string

	list := m.list.View()
	value = m.searchValue

	if m.state == searchState {
		state = m.textinput.View()
	}

	return appStyle.Render(lipgloss.JoinVertical(lipgloss.Left, lipgloss.JoinHorizontal(lipgloss.Top, list, value), state))
}
