package player

import (
	"fmt"
	"log"
	"os"
	"slices"
)

type Player struct {
	Name string
	Sign string
}

var usedSigns = []string{}

func CreatePlayer() (Player, error) {
	name, err := getName()
	if err != nil {
		return Player{}, err
	}

	sign, err := getSign()
	if err != nil {
		return Player{}, err
	}

	return Player{Name: name, Sign: sign}, nil
}

func getName() (string, error) {
	var name string

	_, err := fmt.Scan(&name)

	if err != nil {
		log.Panicln("Ошибка при выборе из меню ", err)
	}

	return name, nil
}

func getSign() (string, error) {
	var sign, choice string
	var err error

	fmt.Printf("Крестик или нолик?\n1. Крестик (X)\n2. Нолик (O)\n3. Выйти(\n")
	_, err = fmt.Scan(&choice)

	if err != nil {
		log.Panicln("Ошибка при выборе из меню ", err)
	}

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
		err = fmt.Errorf("что-то пошло не так")
	}
	return sign, err
}
