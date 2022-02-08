// @description:
// @file: tui.go
// @date: 2022/02/07

// Package tui .
package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	keyMap    KeyMap
	textinput textinput.Model

	selected map[int]struct{}
	choices  []string
	cursor   int

	state
}

func New() *model {
	t := textinput.New()
	t.Prompt = "> "
	t.CharLimit = 256
	t.PlaceholderStyle = lipgloss.NewStyle()

	return &model{
		keyMap:    DefaultKeyMap(),
		textinput: t,

		choices:  []string{"🍎苹果", "🍌香蕉", "🍉西瓜"},
		selected: make(map[int]struct{}),
	}
}
