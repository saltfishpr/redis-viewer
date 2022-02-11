// @description:
// @file: view.go
// @date: 2022/02/08

package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

var (
	appStyle      = lipgloss.NewStyle().Padding(0)
	dividerStyle  = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#9B9B9B", Dark: "#5C5C5C"})
	viewportStyle = lipgloss.NewStyle().Border(lipgloss.HiddenBorder(), false, true)
)

func (m *model) View() string {
	builder := &strings.Builder{}
	divider := dividerStyle.Render(strings.Repeat("-", m.viewport.Width))
	if it := m.list.SelectedItem(); it != nil {
		key := fmt.Sprintf("Key: \n%s\n", it.(item).key)
		keyType := fmt.Sprintf("KeyType: %s\n", it.(item).keyType)
		value := fmt.Sprintf("Value: \n%s\n", it.(item).val)
		builder.WriteString(key)
		builder.WriteString(divider)
		builder.WriteByte('\n')
		builder.WriteString(keyType)
		builder.WriteString(divider)
		builder.WriteByte('\n')
		builder.WriteString(value)
	} else {
		builder.WriteString("No item selected")
	}
	valueDetail := builder.String()
	m.viewport.SetContent(wordwrap.String(valueDetail, m.viewport.Width))

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
