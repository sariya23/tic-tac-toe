package board

import (
	"fmt"
	"strings"
)

const (
	Xsign     = "X"
	Osign     = "O"
	EmptySign = "*"
)

var (
	markColumns = fmt.Sprint(strings.Repeat(" ", 3), "A", strings.Repeat(" ", 4), "B", strings.Repeat(" ", 4), "C")
	outline     = fmt.Sprint(" " + strings.Repeat("-", 15))
)

type Board struct {
	Board [][]string
	SizeX int
	SizeY int
}

type Step struct {
	X int
	Y int
}

func (b *Board) NewBoard() Board {
	emptyRow := strings.Split(strings.Repeat(EmptySign, 3), "")
	for i := 0; i < 3; i++ {
		b.Board = append(b.Board, emptyRow)
	}
	b.SizeX = 3
	b.SizeY = 3
	return *b
}

func (b *Board) DrawBoard() {
	fmt.Println(markColumns)
	fmt.Println(outline)

	for i := 0; i < b.SizeX; i++ {
		row := fmt.Sprint(i + 1)
		for j := 0; j < b.SizeY; j++ {
			row += fmt.Sprintf("| %s |", b.Board[i][j])
		}
		fmt.Println(row)
		fmt.Println(outline)
	}

}

func (b *Board) IsGameEnd() bool {
	var countEmptyFields int

	for i := 0; i < b.SizeX; i++ {
		for j := 0; j < b.SizeY; j++ {
			if b.Board[i][j] == "*" {
				countEmptyFields++
			}
		}
	}

	if countEmptyFields > 0 {
		return false
	}

	return true
}

func (b *Board) GetAvailableSteps() []Step {
	availableSteps := make([]Step, 0)
	for i := 0; i < b.SizeX; i++ {
		for j := 0; j < b.SizeY; j++ {
			if strings.Contains(Xsign+Osign, b.Board[i][j]) {
				availableSteps = append(availableSteps, Step{X: i, Y: j})
			}
		}
	}

	return availableSteps
}
