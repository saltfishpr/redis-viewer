// @file: view.go
// @date: 2022/02/08

package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

var (
	listViewStyle = lipgloss.NewStyle().
			MarginRight(2).
			Border(lipgloss.RoundedBorder(), false, true, false, false)
	dividerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#9B9B9B", Dark: "#5C5C5C"})

	statusNugget   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFDF5")).Padding(0, 1)
	statusBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
			Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})
	statusStyle = lipgloss.NewStyle().
			Inherit(statusBarStyle).
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#FF5F87")).
			Padding(0, 1).
			MarginRight(1)
	encodingStyle = statusNugget.Copy().Background(lipgloss.Color("#A550DF")).Align(lipgloss.Right)
	statusText    = lipgloss.NewStyle().Inherit(statusBarStyle)
	datetimeStyle = statusNugget.Copy().Background(lipgloss.Color("#6124DF"))
)

func (m model) listView() string {
	return listViewStyle.Render(m.list.View())
}

func (m model) detailView() string {
	builder := &strings.Builder{}
	divider := dividerStyle.Render(strings.Repeat("-", m.viewport.Width)) + "\n"
	if it := m.list.SelectedItem(); it != nil {
		key := fmt.Sprintf("Key: \n%s\n", it.(item).key)
		keyType := fmt.Sprintf("KeyType: %s\n", it.(item).keyType)
		value := fmt.Sprintf("Value: \n%s\n", it.(item).val)
		builder.WriteString(key)
		builder.WriteString(divider)
		builder.WriteString(keyType)
		builder.WriteString(divider)
		builder.WriteString(value)
	} else {
		builder.WriteString("No item selected")
	}
	m.viewport.SetContent(wordwrap.String(builder.String(), m.viewport.Width))
	return m.viewport.View()
}

func (m model) statusView() string {
	var status string
	var statusDesc string
	switch m.state {
	case searchState:
		status = "Search"
		statusDesc = m.textinput.View()
	default:
		status = "Ready"
		statusDesc = m.statusMessage
	}

	statusKey := statusStyle.Render(status)
	encoding := encodingStyle.Render("UTF-8")
	datetime := datetimeStyle.Render(time.Now().Format("2006-01-02 15:04:05"))

	statusVal := statusText.Copy().
		Width(m.width - lipgloss.Width(statusKey) - lipgloss.Width(encoding) - lipgloss.Width(datetime)).
		Render(statusDesc)

	bar := lipgloss.JoinHorizontal(lipgloss.Top, statusKey, statusVal, encoding, datetime)

	return statusBarStyle.Width(m.width).Render(bar)
}

func (m model) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.JoinHorizontal(lipgloss.Top, m.listView(), m.detailView()),
		m.statusView(),
	)
}
