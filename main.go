package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"tic-tac-toe/board"
	"tic-tac-toe/player"
)

func main() {
	var isGameEnd, isDraw bool
	var b board.Board

	SetLog()
	fmt.Println("Введите имя первого игрока: ")
	player1, err := player.CreatePlayer()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("Введите имя второго игрока: ")
	player2, err := player.CreatePlayer()
	if err != nil {
		log.Panicln(err)
	}

	log.Println(player1, player2)

	fmt.Printf("%s - %s VS %s - %s\n", player1.Name, player1.Sign, player2.Name, player2.Sign)

	player.WhoStepFirst(&player1, &player2)

	log.Println(player1, player2)

	b = b.NewBoard()
	log.Println(b)
	fmt.Println("ИГРА НАЧИНАЕТСЯ!!!")
	fmt.Println()

	for {
		isGameEnd, isDraw = b.IsGameEnd()

		if isGameEnd {
			break
		}

		player.PrintWhoStep(player1, player2)

		b.DrawBoard()
		b.ShowAvailablaSteps()

		stepChoice := GetInput()
		choiceNumber, err := ConverToNumber(stepChoice)
		availableSteps := b.GetAvailableSteps()

		if err != nil || choiceNumber <= 0 || choiceNumber > len(availableSteps) {
			for {
				fmt.Println("Нужно ввести номер из списка")
				stepChoice = GetInput()
				choiceNumber, err = ConverToNumber(stepChoice)
				if err != nil || choiceNumber <= 0 || choiceNumber > len(availableSteps) {
					continue
				} else {
					break
				}
			}
		}

		translatedCoord := board.TranslateStepToLetter(availableSteps[choiceNumber-1])

		if player1.IsStep {
			fmt.Printf("Игрок \"%s\" поставил(а) \"%s\" на поле %s\n", player1.Name, player1.Sign, translatedCoord.X+translatedCoord.Y)
			b.MarkStep(availableSteps[choiceNumber-1], player1.Sign)
		} else {
			fmt.Printf("Игрок \"%s\" поставил(а) \"%s\" на поле %s\n", player2.Name, player2.Sign, translatedCoord.X+translatedCoord.Y)
			b.MarkStep(availableSteps[choiceNumber-1], player2.Sign)
		}

		player1.IsStep = !player1.IsStep
		player2.IsStep = !player2.IsStep
	}

	fmt.Println()

	if isDraw {
		fmt.Println("Игра заканчивается нечьей")
	} else if player2.IsStep {
		fmt.Printf("Выигрывает игрок %s - %s\n", player2.Name, player2.Sign)
	} else {
		fmt.Printf("Выигрывает игрок %s - %s\n", player1.Name, player1.Sign)
	}

	b.DrawBoard()
}

func SetLog() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	log.SetOutput(file)
}

func GetInput() string {
	reader := bufio.NewReader(os.Stdin)
	choice, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal("Что-то пошло не так при вводе")
	}

	return strings.Trim(choice, "\n")
}

func ConverToNumber(s string) (int, error) {
	n, err := strconv.Atoi(s)

	if err != nil {
		log.Println(err)
		return 0, err
	}

	return n, nil
}
