package main

import (
	"fmt"
	"os"
	"slices"
)

type Player struct {
	Name string
	Sign string
}

var usedSigns = []string{}

func main() {
	fmt.Println("Введите данные по первому игроку: ")
	player1, err := createPlayer()
	if err != nil {
		panic(err)
	}

	fmt.Println("Введите данные по второму игроку: ")
	player2, err := createPlayer()
	if err != nil {
		panic(err)
	}
	fmt.Println(player1, player2)

}

func createPlayer() (Player, error) {
	var choice, username string

	_, err := fmt.Scan(&username)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Крестик или нолик?\n1. Крестик (X)\n2. Нолик (O)\n3. Выйти(\n")
	_, err = fmt.Scan(&choice)
	if err != nil {
		panic(err)
	}

	sign, err := getSign(choice)
	if err != nil {
		panic(err)
	}

	return Player{Name: username, Sign: sign}, nil
}

func getSign(choice string) (string, error) {
	var sign string
	var err error

	if choice == "1" && !slices.Contains(usedSigns, "X") {
		usedSigns = append(usedSigns, "X")
		sign = "X"
	} else if choice == "2" && !slices.Contains(usedSigns, "O") {
		usedSigns = append(usedSigns, "O")
		sign = "O"
	} else if choice == "3" {
		fmt.Println("Бай-бай")
		sign = ""
		os.Exit(0)
	} else {
		err = fmt.Errorf("Что-то пошло не так...")
	}
	return sign, err
}
