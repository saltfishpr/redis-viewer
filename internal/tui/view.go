// @description:
// @file: view.go
// @date: 2022/02/08

package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

var (
	appStyle = lipgloss.NewStyle().Padding(1, 2)

	viewportStyle = lipgloss.NewStyle().Border(lipgloss.HiddenBorder(), false, false, false, true)
)

func (m *model) View() string {
	if it := m.list.SelectedItem(); it != nil {
		valueDetail := fmt.Sprintf("KeyType: %s\nValue:\n%s", it.(item).keyType, it.(item).val)
		m.viewport.SetContent(wordwrap.String(valueDetail, m.viewport.Width))
	}

	if m.state == searchState {
		m.stateDesc = m.textinput.View()
	}

	return appStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			lipgloss.JoinHorizontal(lipgloss.Top, m.list.View(), viewportStyle.Render(m.viewport.View())),
			m.stateDesc,
		),
	)
}
