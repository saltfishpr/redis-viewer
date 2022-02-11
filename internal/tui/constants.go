// @description:
// @file: constants.go
// @date: 2022/02/10

package tui

type state int

const (
	defaultState state = iota
	searchState
)

const (
	listProportion  = 0.3
	statusBarHeight = 1
	maxScanCount    = 9999
)
