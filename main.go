package main

import (
	"fmt"
	"log"
	"os"

	"tic-tac-toe/board"
	"tic-tac-toe/player"
)

func main() {
	var b board.Board

	SetLog()
	fmt.Println("Введите данные по первому игроку: ")
	player1, err := player.CreatePlayer()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("Введите данные по второму игроку: ")
	player2, err := player.CreatePlayer()
	if err != nil {
		log.Panicln(err)
	}

	log.Println(player1, player2)

	fmt.Printf("%s - %s VS %s - %s\n", player1.Name, player1.Sign, player2.Name, player2.Sign)

	player1.IsStep = player.RandomBool()
	player2.IsStep = !player1.IsStep

	log.Println(player1, player2)

	if player1.IsStep {
		fmt.Printf("Первым ходит игрок %s\n", player1.Name)
	} else {
		fmt.Printf("Первым ходит игрок %s\n", player2.Name)
	}

	b = b.NewBoard()
	log.Println(b)
	b.DrawBoard()

}

func SetLog() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	log.SetOutput(file)
}
