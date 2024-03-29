package board

import "testing"

type testBoardPairs struct {
	board     Board
	isGameEnd bool
}

var testsIsGameEnd = []testBoardPairs{
	{board: Board{Board: [][]string{{"X", "X", "X"}, {"O", "X", "O"}, {"X", "O", "X"}}, SizeX: 3, SizeY: 3}, isGameEnd: true},
	{board: Board{Board: [][]string{{"*", "X", "X"}, {"O", "X", "O"}, {"X", "O", "X"}}, SizeX: 3, SizeY: 3}, isGameEnd: false},
	{board: Board{Board: [][]string{{"*", "*", "*"}, {"*", "*", "*"}, {"*", "*", "*"}}, SizeX: 3, SizeY: 3}, isGameEnd: false},
	{board: Board{Board: [][]string{{"X", "*", "*"}, {"*", "*", "*"}, {"*", "*", "*"}}, SizeX: 3, SizeY: 3}, isGameEnd: false},
}

func TestIsGameEnd(t *testing.T) {
	for _, pair := range testsIsGameEnd {
		v := pair.board.IsGameEnd()

		if v != pair.isGameEnd {
			t.Error(
				"For", pair.board,
				"expected", pair.isGameEnd,
				"got", v,
			)
		}
	}
}
