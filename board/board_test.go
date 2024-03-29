package board

import (
	"reflect"
	"testing"
)

type testStructIsGameEnd struct {
	board             Board
	expectedIsGameEnd bool
}

type testStructGetAvailableSteps struct {
	board         Board
	expectedSteps []Step
}

var expectedNewBoard = [][]string{
	{"*", "*", "*"},
	{"*", "*", "*"},
	{"*", "*", "*"},
}

var testsIsGameEnd = []testStructIsGameEnd{
	{board: Board{Board: [][]string{{"X", "X", "X"}, {"O", "X", "O"}, {"X", "O", "X"}}, SizeX: 3, SizeY: 3}, expectedIsGameEnd: true},
	{board: Board{Board: [][]string{{"*", "X", "X"}, {"O", "X", "O"}, {"X", "O", "X"}}, SizeX: 3, SizeY: 3}, expectedIsGameEnd: false},
	{board: Board{Board: [][]string{{"*", "*", "*"}, {"*", "*", "*"}, {"*", "*", "*"}}, SizeX: 3, SizeY: 3}, expectedIsGameEnd: false},
	{board: Board{Board: [][]string{{"X", "*", "*"}, {"*", "*", "*"}, {"*", "*", "*"}}, SizeX: 3, SizeY: 3}, expectedIsGameEnd: false},
}

var testsGetAvailableSteps = []testStructGetAvailableSteps{
	{
		board: Board{
			Board: [][]string{{"*", "*", "*"}, {"*", "*", "*"}, {"*", "*", "*"}},
			SizeX: 3,
			SizeY: 3,
		},
		expectedSteps: []Step{
			{0, 0},
			{0, 1},
			{0, 2},
			{1, 0},
			{1, 1},
			{1, 2},
			{2, 0},
			{2, 1},
			{2, 2},
		},
	},
	{
		board: Board{
			Board: [][]string{{"X", "X", "X"}, {"O", "X", "O"}, {"X", "O", "X"}},
			SizeX: 3,
			SizeY: 3,
		},
		expectedSteps: []Step{},
	},
	{
		board: Board{
			Board: [][]string{
				{"X", "X", "X"},
				{"O", "*", "O"},
				{"X", "O", "*"},
			},
			SizeX: 3,
			SizeY: 3,
		},
		expectedSteps: []Step{
			{1, 1},
			{2, 2},
		},
	},
}

func TestNewBoard(t *testing.T) {
	var b Board
	b = b.NewBoard()
	actualBoard := b.Board

	if !reflect.DeepEqual(actualBoard, expectedNewBoard) {
		t.Error("New board created wrong")
	}
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

func TestGetAvailableSteps(t *testing.T) {
	for _, pair := range testsGetAvailableSteps {
		actual := pair.board.GetAvailableSteps()
		if !reflect.DeepEqual(actual, pair.expectedSteps) {
			t.Error(
				"For", pair.board,
				"expected", pair.expectedSteps,
				"got", actual,
			)
		}
	}
}
