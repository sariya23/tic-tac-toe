package board

import (
	"fmt"
	"strings"
	"tic-tac-toe/player"
)

type Board struct {
	Board [][]string
	SizeX int
	SizeY int
}

func (b *Board) NewBoard() Board {
	for i := 0; i < 3; i++ {
		b.Board = append(b.Board, []string{"*", "*", "*"})
	}
	b.SizeX = 3
	b.SizeY = 3
	return *b
}

func (b *Board) DrawBoard() {
	fmt.Println(strings.Repeat("-", 15))

	for i := 0; i < b.SizeX; i++ {
		row := ""
		for j := 0; j < b.SizeY; j++ {
			row += fmt.Sprintf("| %s |", b.Board[i][j])
		}
		fmt.Println(row)
		fmt.Println(strings.Repeat("-", 15))
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

func (b *Board) GetAvailableSteps(p *player.Player) [][]int {
	availableSteps := make([][]int, 0)
	for i := 0; i < b.SizeX; i++ {
		for j := 0; j < b.SizeY; j++ {
			if b.Board[i][j] != p.Sign {
				availableSteps = append(availableSteps, []int{i, j})
			}
		}
	}

	return availableSteps
}
