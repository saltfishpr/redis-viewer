// @description:
// @file: command.go
// @date: 2022/02/07

package tui

import (
	"context"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type scanMsg struct {
	items []list.Item
}

func (m model) scanCmd() tea.Cmd {
	return func() tea.Msg {
		ctx := context.Background()
		iter := m.rdb.Scan(ctx, 0, m.searchValue, defaultScanCount).Iterator()
		var items []list.Item
		for iter.Next(ctx) {
			key := iter.Val()
			val, err := m.rdb.Get(ctx, key).Result()
			if err != nil {
				items = append(items, item{key: key, val: err.Error(), err: true})
			} else {
				items = append(items, item{key: key, val: val, err: false})
			}
		}
		if err := iter.Err(); err != nil {
			m.stateDesc = err.Error()
			return nil
		}
		return scanMsg{items: items}
	}
}
