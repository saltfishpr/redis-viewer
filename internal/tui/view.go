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
	listViewStyle = lipgloss.NewStyle().
			PaddingRight(1).
			MarginRight(1).
			Border(lipgloss.RoundedBorder(), false, true, false, false)
	dividerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#9B9B9B", Dark: "#5C5C5C"})

	statusNugget   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFDF5")).Padding(0, 1)
	statusBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
			Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})
	statusStyle = statusBarStyle.
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#FF5F87")).
			Padding(0, 1).
			MarginRight(1)
	pageStyle     = statusNugget.Background(lipgloss.Color("#2D8CFF"))
	encodingStyle = statusNugget.Background(lipgloss.Color("#A550DF")).Align(lipgloss.Right)
	statusText    = statusBarStyle
	datetimeStyle = statusNugget.Background(lipgloss.Color("#6124DF"))
)

func (m model) listView() string {
	listModel := m.list
	total := len(listModel.Items())
	current := 0
	if listModel.SelectedItem() != nil {
		current = listModel.GlobalIndex() + 1
	}
	listModel.Title = fmt.Sprintf("Redis Viewer (%d/%d)", current, total)

	return listViewStyle.Render(listModel.View())
}

func (m model) viewportContent(width int) string {
	var builder strings.Builder
	divider := dividerStyle.Render(strings.Repeat("-", width)) + "\n"
	if it := m.list.SelectedItem(); it != nil {
		keyType := fmt.Sprintf("Type: %s\n", it.(item).keyType)
		key := fmt.Sprintf("Key: \n%s\n", it.(item).key)
		value := fmt.Sprintf("Value: \n%s\n", it.(item).val)
		builder.WriteString(keyType)
		builder.WriteString(divider)
		builder.WriteString(key)
		builder.WriteString(divider)
		builder.WriteString(value)
	} else {
		builder.WriteString("No item selected")
	}

	return wordwrap.String(builder.String(), width)
}

func (m model) detailView() string {
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
		if !m.countReady {
			status = m.spinner.View()
			statusDesc = "Loading..."
		}
	}

	statusKey := statusStyle.Render(status)
	page := pageStyle.Render(fmt.Sprintf("Page %d Size %d", m.pageNum+1, m.pageSize))
	encoding := encodingStyle.Render("UTF-8")
	datetime := datetimeStyle.Render(m.now)

	statusVal := statusText.
		Width(m.width - lipgloss.Width(statusKey) - lipgloss.Width(page) - lipgloss.Width(encoding) - lipgloss.Width(datetime)).
		Render(statusDesc)

	bar := lipgloss.JoinHorizontal(lipgloss.Top, statusKey, statusVal, page, encoding, datetime)

	return statusBarStyle.Width(m.width).Render(bar)
}

func (m model) View() string {
	// TODO: refresh status view only
	return lipgloss.JoinVertical(lipgloss.Left,
		lipgloss.JoinHorizontal(lipgloss.Top, m.listView(), m.detailView()),
		m.statusView(),
	)
}
