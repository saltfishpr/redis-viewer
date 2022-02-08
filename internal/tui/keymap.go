// @description: 快捷键映射
// @file: keymap.go
// @date: 2022/02/07

package tui

import "github.com/charmbracelet/bubbles/key"

// KeyMap defines the keybindings for the app.
type KeyMap struct {
	Quit   key.Binding
	Down   key.Binding
	Up     key.Binding
	Left   key.Binding
	Right  key.Binding
	Space  key.Binding
	Insert key.Binding
	Enter  key.Binding
	Escape key.Binding
}

// DefaultKeyMap returns a set of default keybindings.
func DefaultKeyMap() KeyMap {
	return KeyMap{
		Quit: key.NewBinding(
			key.WithKeys("ctrl+c", "q"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j"),
		),
		Up: key.NewBinding(
			key.WithKeys("up", "k"),
		),
		Left: key.NewBinding(
			key.WithKeys("left", "h"),
		),
		Right: key.NewBinding(
			key.WithKeys("right", "l"),
		),
		Space: key.NewBinding(
			key.WithKeys(" "),
		),
		Insert: key.NewBinding(
			key.WithKeys("i"),
		),
		Enter: key.NewBinding(
			key.WithKeys("enter"),
		),
		Escape: key.NewBinding(
			key.WithKeys("esc"),
		),
	}
}
