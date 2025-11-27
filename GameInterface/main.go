package main

import (
	"fmt"
)

// implementing 2 different console games using interface:

type ConsoleGame interface {
	displayGameRules()
	runGame()
	endGame()
}

type Connect3 ConsoleGame 

type Sudoku ConsoleGame

func main () {
	var choice int
	fmt.Println("Type\n 1 for connect 3\n 2 for sudoku")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Running Connect 3.....")
	case 2:
		fmt.Println("Running Sudoku......")
	}
}