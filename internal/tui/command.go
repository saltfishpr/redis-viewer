// @description:
// @file: command.go
// @date: 2022/02/07

package tui

import (
	"context"
	"fmt"
	"redis-viewer/internal/util"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type scanMsg struct {
	items []list.Item
}

func (m model) scanCmd() tea.Cmd {
	return func() tea.Msg {
		ctx := context.Background()
		var (
			val   interface{}
			err   error
			items []list.Item
		)

		iter := m.rdb.Scan(ctx, 0, m.searchValue, defaultScanCount).Iterator()
		for iter.Next(ctx) {
			key := iter.Val()
			kt := m.rdb.Type(ctx, key).Val()
			switch kt {
			case "string":
				val, err = m.rdb.Get(ctx, key).Result()
			case "list":
				val, err = m.rdb.LRange(ctx, key, 0, -1).Result()
			case "set":
				val, err = m.rdb.SMembers(ctx, key).Result()
			case "zset":
				val, err = m.rdb.ZRange(ctx, key, 0, -1).Result()
			case "hash":
				val, err = m.rdb.HGetAll(ctx, key).Result()
			default:
				val = ""
				err = fmt.Errorf("unsupported type: %s", kt)
			}
			if err != nil {
				items = append(items, item{keyType: kt, key: key, val: err.Error(), err: true})
			} else {
				valBts, _ := util.JsonMarshal(val)
				items = append(items, item{keyType: kt, key: key, val: string(valBts)})
			}
		}
		if err := iter.Err(); err != nil {
			m.stateDesc = err.Error()
			return nil
		}
		return scanMsg{items: items}
	}
}
