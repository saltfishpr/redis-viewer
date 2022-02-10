// @description:
// @file: tui.go
// @date: 2022/02/07

// Package tui .
package tui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

type state int

const (
	defaultState state = iota
	searchState
)

type model struct {
	list      list.Model
	textinput textinput.Model

	searchValue string

	keyMap

	state
}

func New() *model {
	t := textinput.New()
	t.Prompt = "> "
	t.CharLimit = 256
	t.Placeholder = "Enter new item"
	t.PlaceholderStyle = lipgloss.NewStyle()

	items := []list.Item{
		item{title: "Apple", description: "🍎"},
		item{title: "Banana", description: "🍌"},
		item{title: "Cherry", description: "🍒"},
		item{title: "Date", description: "🍅"},
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Which fruit do you want?"
	l.SetShowHelp(false)

	return &model{
		list:      l,
		textinput: t,

		searchValue: "[Value]",

		keyMap: defaultKeyMap(),

		state: defaultState,
	}
}
