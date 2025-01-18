package entity

import (
	"snacke/app/game/setting"
)

type Board struct {
}

func (b *Board) Create(setting setting.Settings) [][]rune {
	board := make([][]rune, setting.Height)
	for i := range board {
		board[i] = make([]rune, setting.Width)
		for j := range board[i] {
			b.fillEdge(i, j, board, setting)
		}
	}

	return board
}

func (b *Board) fillEdge(i, j int, board [][]rune, setting setting.Settings) {
	if i == 0 || j == 0 || i == setting.Height-1 || j == setting.Width-1 {
		board[i][j] = '.'

		return
	}

	board[i][j] = ' '
}
