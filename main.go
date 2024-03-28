package main

import (
	"fmt"
	"log"
	"os"

	"tic-tac-toe/player"
)

func main() {
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

}

func SetLog() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	log.SetOutput(file)
}
