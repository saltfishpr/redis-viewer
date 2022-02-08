// @description:
// @file: view.go
// @date: 2022/02/08

package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m *model) View() string {
	s := "What should we buy at the market?\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"

	input := m.textinput.View()

	if m.state == insertState {
		return lipgloss.JoinVertical(lipgloss.Left, s, input)
	}
	return lipgloss.JoinVertical(lipgloss.Left, s)
}
