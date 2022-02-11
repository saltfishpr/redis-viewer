// @description: 定义常量
// @file: constants.go
// @date: 2022/02/10

package tui

type state int

const (
	defaultState state = iota
	searchState
)

const (
	maxScanCount = 9999

	listProportion = 0.3
)
