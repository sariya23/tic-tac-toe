package board

import (
	"fmt"
	"strings"
)

type Board struct {
	Board [][]string
}

func (b *Board) NewBoard() Board {
	for i := 0; i < 3; i++ {
		b.Board = append(b.Board, []string{"*", "*", "*"})
	}

	return *b
}

func (b *Board) DrawBoard() {
	fmt.Println(strings.Repeat("-", 15))

	for i := 0; i < len(b.Board); i++ {
		row := ""
		for j := 0; j < len(b.Board[i]); j++ {
			row += fmt.Sprintf("| %s |", b.Board[i][j])
		}
		fmt.Println(row)
		fmt.Println(strings.Repeat("-", 15))
	}

}
