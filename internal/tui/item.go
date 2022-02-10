// @description: list item
// @file: item.go
// @date: 2022/02/08

package tui

import "strconv"

type item struct {
	key string
	val string

	err bool
}

func (i item) Title() string { return i.key }
func (i item) Description() string {
	if i.err {
		return "get error: " + i.val
	}
	return "keySize: " + strconv.Itoa(len(i.key)) + ", valueSize: " + strconv.Itoa(len(i.val))
}
func (i item) FilterValue() string { return i.key }
