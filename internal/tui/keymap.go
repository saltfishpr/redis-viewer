// @description: 快捷键映射
// @file: keymap.go
// @date: 2022/02/07

package tui

import "github.com/charmbracelet/bubbles/key"

// keyMap defines the keybindings for the app.
type keyMap struct {
	search key.Binding
}

// defaultKeyMap returns a set of default keybindings.
func defaultKeyMap() keyMap {
	return keyMap{
		search: key.NewBinding(
			key.WithKeys("s"),
		),
	}
}
