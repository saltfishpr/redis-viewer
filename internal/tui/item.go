// @description: list item
// @file: item.go
// @date: 2022/02/08

package tui

type item struct {
	title       string
	description string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.description }
func (i item) FilterValue() string { return i.title }
