// @file: tui.go
// @date: 2022/02/07

// Package tui .
package tui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
	"github.com/go-redis/redis/v8"
)

type model struct {
	width, height int

	list      list.Model
	textinput textinput.Model
	viewport  viewport.Model
	spinner   spinner.Model

	rdb           redis.Cmdable
	searchValue   string
	statusMessage string
	ready         bool

	keyMap
	state
}

func New(rdb redis.Cmdable) *model {
	t := textinput.New()
	t.Prompt = "> "
	t.Placeholder = "Search Key"
	t.PlaceholderStyle = lipgloss.NewStyle()

	l := list.New(nil, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Redis Viewer by SaltFishPr"
	l.SetShowFilter(false)
	l.SetShowHelp(false)
	l.SetFilteringEnabled(false)

	s := spinner.New()
	s.Spinner = spinner.Dot

	return &model{
		list:      l,
		textinput: t,
		spinner:   s,

		rdb: rdb,

		keyMap: defaultKeyMap(),
		state:  defaultState,
	}
}
