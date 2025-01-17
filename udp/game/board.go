package game

type Board struct {
}

func (b *Board) Create() [][]rune {
	board := make([][]rune, Setting.Height)
	for i := range board {
		board[i] = make([]rune, Setting.Width)
		for j := range board[i] {
			b.fillEdge(i, j, board)
		}
	}

	return board
}

func (b *Board) fillEdge(i, j int, board [][]rune) {
	if i == 0 || j == 0 || i == Setting.Height-1 || j == Setting.Width-1 {
		board[i][j] = '.'

		return
	}

	board[i][j] = ' '
}
