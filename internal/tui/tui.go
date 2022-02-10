// @description:
// @file: tui.go
// @date: 2022/02/07

// Package tui .
package tui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
	"github.com/go-redis/redis/v8"
)

type model struct {
	list      list.Model
	textinput textinput.Model

	rdb         *redis.Client
	searchValue string
	valueDetail string
	stateDesc   string

	keyMap
	state
}

func New(rdb *redis.Client) *model {
	t := textinput.New()
	t.Prompt = "> "
	t.Placeholder = "Search Key"
	t.PlaceholderStyle = lipgloss.NewStyle()

	l := list.New(nil, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Redis Viewer by SaltFishPr"
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)

	return &model{
		list:      l,
		textinput: t,

		rdb: rdb,

		keyMap: defaultKeyMap(),
		state:  defaultState,
	}
}
