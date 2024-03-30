package board

import (
	"reflect"
	"testing"
)

type testStructIsGameEnd struct {
	board             Board
	expectedIsGameEnd bool
	expectedIsDraw    bool
}

type testStructIsWinningValues struct {
	values            []string
	expectedIsWinning bool
}

type testStructGetAvailableSteps struct {
	board         Board
	expectedSteps []stepCoordinates
}

type testStructGetRowValues struct {
	board             Board
	rowIndex          int
	expectedRowValues []string
}

type testStructGetColValues struct {
	board             Board
	colIndex          int
	expectedColValues []string
}

type testStructTranslateCoordinates struct {
	coordinate           stepCoordinates
	translatedCoordinate StepLetterCoordinates
}

var expectedNewBoard = [][]string{
	{"*", "*", "*"},
	{"*", "*", "*"},
	{"*", "*", "*"},
}

var testsIsGameEnd = []testStructIsGameEnd{
	{board: Board{Board: [][]string{{"X", "X", "X"}, {"O", "X", "O"}, {"X", "O", "X"}}, SizeX: 3, SizeY: 3}, expectedIsGameEnd: true, expectedIsDraw: false},
	{board: Board{Board: [][]string{{"*", "X", "X"}, {"O", "X", "O"}, {"X", "O", "X"}}, SizeX: 3, SizeY: 3}, expectedIsGameEnd: false, expectedIsDraw: false},
	{board: Board{Board: [][]string{{"*", "*", "*"}, {"*", "*", "*"}, {"*", "*", "*"}}, SizeX: 3, SizeY: 3}, expectedIsGameEnd: false, expectedIsDraw: false},
	{board: Board{Board: [][]string{{"X", "*", "*"}, {"*", "*", "*"}, {"*", "*", "*"}}, SizeX: 3, SizeY: 3}, expectedIsGameEnd: false, expectedIsDraw: false},
	{board: Board{Board: [][]string{{"X", "*", "O"}, {"O", "X", "*"}, {"*", "*", "X"}}, SizeX: 3, SizeY: 3}, expectedIsGameEnd: true, expectedIsDraw: false},
	{board: Board{Board: [][]string{{"O", "X", "*"}, {"O", "X", "X"}, {"O", "*", "X"}}, SizeX: 3, SizeY: 3}, expectedIsGameEnd: true, expectedIsDraw: false},
}

var testsIsWinningValues = []testStructIsWinningValues{
	{values: []string{"X", "X", "X"}, expectedIsWinning: true},
	{values: []string{"O", "O", "O"}, expectedIsWinning: true},
	{values: []string{"X", "O", "X"}, expectedIsWinning: false},
	{values: []string{"*", "*", "*"}, expectedIsWinning: false},
}

var testsTranslateCoordinate = []testStructTranslateCoordinates{
	{coordinate: stepCoordinates{X: 1, Y: 2}, translatedCoordinate: StepLetterCoordinates{X: "2", Y: "C"}},
	{coordinate: stepCoordinates{X: 0, Y: 0}, translatedCoordinate: StepLetterCoordinates{X: "1", Y: "A"}},
	{coordinate: stepCoordinates{X: 2, Y: 2}, translatedCoordinate: StepLetterCoordinates{X: "3", Y: "C"}},
}

var testsGetColValues = []testStructGetColValues{
	{board: Board{Board: [][]string{{"X", "X", "X"}, {"O", "X", "O"}, {"X", "O", "X"}}, SizeX: 3, SizeY: 3}, colIndex: 0, expectedColValues: []string{"X", "O", "X"}},
	{board: Board{Board: [][]string{{"X", "X", "X"}, {"O", "X", "O"}, {"X", "O", "X"}}, SizeX: 3, SizeY: 3}, colIndex: 2, expectedColValues: []string{"X", "O", "X"}},
}

var testsGetRowValues = []testStructGetRowValues{
	{board: Board{Board: [][]string{{"X", "X", "X"}, {"O", "X", "O"}, {"X", "O", "X"}}, SizeX: 3, SizeY: 3}, rowIndex: 0, expectedRowValues: []string{"X", "X", "X"}},
	{board: Board{Board: [][]string{{"X", "X", "X"}, {"O", "X", "O"}, {"X", "O", "X"}}, SizeX: 3, SizeY: 3}, rowIndex: 2, expectedRowValues: []string{"X", "O", "X"}},
}

var testsGetAvailableSteps = []testStructGetAvailableSteps{
	{
		board: Board{
			Board: [][]string{{"*", "*", "*"}, {"*", "*", "*"}, {"*", "*", "*"}},
			SizeX: 3,
			SizeY: 3,
		},
		expectedSteps: []stepCoordinates{
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
		expectedSteps: []stepCoordinates{},
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
		expectedSteps: []stepCoordinates{
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

func TestWinningValues(t *testing.T) {
	for _, pair := range testsIsWinningValues {
		actualIsWinnigValues := isWinningValues(pair.values)

		if actualIsWinnigValues != pair.expectedIsWinning {
			t.Error(
				"For", pair,
				"expected", pair.expectedIsWinning,
				"got", actualIsWinnigValues,
			)
		}
	}
}

func TestGetRowValues(t *testing.T) {
	for _, pair := range testsGetRowValues {
		actualRowValues := pair.board.getRowValues(pair.rowIndex)

		if !reflect.DeepEqual(actualRowValues, pair.expectedRowValues) {
			t.Error(
				"For", pair,
				"expected", pair.expectedRowValues,
				"got", actualRowValues,
			)
		}
	}
}

func TestGetColValues(t *testing.T) {
	for _, pair := range testsGetColValues {
		actualColValues := pair.board.getColValues(pair.colIndex)

		if !reflect.DeepEqual(actualColValues, pair.expectedColValues) {
			t.Error(
				"For", pair,
				"expected", pair.expectedColValues,
				"got", actualColValues,
			)
		}
	}
}

func TestIsGameEnd(t *testing.T) {
	for _, pair := range testsIsGameEnd {
		actualIsGameEnd, actualIsDraw := pair.board.IsGameEnd()

		if actualIsGameEnd != pair.expectedIsGameEnd || actualIsDraw != pair.expectedIsDraw {
			t.Error(
				"For", pair.board,
				"expected", pair.expectedIsGameEnd, pair.expectedIsDraw,
				"got", actualIsGameEnd, actualIsDraw,
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

func TestTranslateCoordinates(t *testing.T) {
	for _, pair := range testsTranslateCoordinate {
		actual := TranslateStepToLetter(pair.coordinate)

		if actual != pair.translatedCoordinate {
			t.Error(
				"For", pair.coordinate,
				"expected", pair.translatedCoordinate,
				"got", actual,
			)
		}
	}
}
