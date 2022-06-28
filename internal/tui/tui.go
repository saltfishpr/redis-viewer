// @file: tui.go
// @date: 2022/02/07

// Package tui .
package tui

import (
	"context"
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
	"github.com/go-redis/redis/v8"
	"github.com/saltfishpr/redis-viewer/internal/conf"
	"github.com/saltfishpr/redis-viewer/internal/constant"
)

type state int

const (
	defaultState state = iota
	searchState
)

//nolint:govet
type model struct {
	width, height int

	list      list.Model
	textinput textinput.Model
	viewport  viewport.Model
	spinner   spinner.Model

	rdb           redis.UniversalClient
	searchValue   string
	statusMessage string
	ready         bool
	now           string

	offset int64
	limit  int64 // scan size

	keyMap
	state
}

func New(config conf.Config) (*model, error) {
	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:        config.Addrs,
		DB:           config.DB,
		Username:     config.Username,
		Password:     config.Password,
		MaxRetries:   constant.MaxRetries,
		MaxRedirects: constant.MaxRedirects,
		MasterName:   config.MasterName,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("connect to redis failed: %w", err)
	}

	t := textinput.New()
	t.Prompt = "> "
	t.Placeholder = "Search Key"
	t.PlaceholderStyle = lipgloss.NewStyle()

	l := list.New(nil, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Redis Viewer"
	l.SetShowHelp(false)
	l.SetShowStatusBar(false)
	l.SetShowFilter(false)
	l.SetFilteringEnabled(false)

	s := spinner.New()
	s.Spinner = spinner.Dot

	return &model{
		list:      l,
		textinput: t,
		spinner:   s,

		rdb: rdb,

		limit: config.Limit,

		keyMap: defaultKeyMap(),
		state:  defaultState,
	}, nil
}
