package board

import "testing"

type testStructIsGameEnd struct {
	board             Board
	expectedIsGameEnd bool
}

var testsIsGameEnd = []testStructIsGameEnd{
	{board: Board{Board: [][]string{{"X", "X", "X"}, {"O", "X", "O"}, {"X", "O", "X"}}, SizeX: 3, SizeY: 3}, expectedIsGameEnd: true},
	{board: Board{Board: [][]string{{"*", "X", "X"}, {"O", "X", "O"}, {"X", "O", "X"}}, SizeX: 3, SizeY: 3}, expectedIsGameEnd: false},
	{board: Board{Board: [][]string{{"*", "*", "*"}, {"*", "*", "*"}, {"*", "*", "*"}}, SizeX: 3, SizeY: 3}, expectedIsGameEnd: false},
	{board: Board{Board: [][]string{{"X", "*", "*"}, {"*", "*", "*"}, {"*", "*", "*"}}, SizeX: 3, SizeY: 3}, expectedIsGameEnd: false},
}

func TestIsGameEnd(t *testing.T) {
	for _, pair := range testsIsGameEnd {
		actual := pair.board.IsGameEnd()

		if actual != pair.expectedIsGameEnd {
			t.Error(
				"For", pair.board,
				"expected", pair.expectedIsGameEnd,
				"got", actual,
			)
		}
	}
}
