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
	if it := m.list.SelectedItem(); it != nil {
		m.valueDetail = "Value:\n" + it.(item).val
	}

	if m.state == searchState {
		m.stateDesc = m.textinput.View()
	}

	return appStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			lipgloss.JoinHorizontal(lipgloss.Top, m.list.View(), m.valueDetail),
			m.stateDesc,
		),
	)
}
