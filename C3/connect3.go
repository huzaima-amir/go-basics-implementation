package main
// implementation of connect 3 game to apply topic 1, 2 and 3 concepts, including slices, structs, loops, string handling, pointers

// Players enter number of rounds -> players enter usernames -> outer loop for game round
//  (empty board prints out, with current scores, players get turn one by one, loop exits if a player scores, or no more spaces left.)
// player making move should trigger applyMove, will need a variable to hold board state throughout round
//  need to figure out a way to print out player turns in order correctly, 
// need inner loop to check 
import (
	"fmt"
	"strings"
)
func createBoard() [][]string {
	// slice of slices to represent game board
    board := make([][]string, 3)
    for i := 0; i < 3; i++ {
        board[i] = make([]string, 3)
        for j := 0; j < 3; j++ {
            board[i][j] = "_"
        }
    }
    return board
}

func displayBoard(board [][]string) { //display board at the start of each round and after each move application
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

type Coordinates struct {
	x, y int
}

func applyMove(P1,P2 Player, count int,board [][] string) { // to apply game move to board
	var currentPlayer Player
	var c Coordinates

	if count/2 == 0{
		currentPlayer = P1
	} else {
		currentPlayer = P2
	}
	fmt.Println(currentPlayer.symbol, "make your move.\n x coordinate:")
	fmt.Scan(&c.x) // x coordinate for input symbol
	fmt.Println("y coordinate:")
	fmt.Scan(&c.y)

	board[c.x][c.y] = currentPlayer.symbol
	checkRoundWinner(board)
}

// Helpers for ending round loop:
func checkRoundWinner(board [][]string) bool { // check for a win to at each iteration of board to increase score and end round loop
	return false
}

func hasEmptySpaces(board [][]string) bool { // to check for empty spaces on board to end the round, if no one scores
    for _, row := range board {
        for _, cell := range row {
            if cell == "_" {
                return true
            }
        }
    }
    return false
}

type Player struct{
	symbol string 
	score int
}

func main(){
	var n int  // amount of game rounds
	P1 := Player{"", 0}
	P2 := Player{"", 0}
	var winner Player 

	defer fmt.Println("Game Over!\n", winner.symbol, "won, score:", winner.score)
	fmt.Println("Enter number of game rounds:")
	fmt.Scan(&n)
	//fmt.Println(createBoard(n))
	fmt.Println("Enter Player 1 Symbol:")
	fmt.Scan(&P1.symbol)
	fmt.Println("Enter Player 2 Symbol:")
	fmt.Scan(&P2.symbol)

	for i := 0; i < n; i++ { //simulating game rounds
		gameBoard := createBoard()
		count := 0
		
		for { //loops forever unless if condition for break is true
			if checkRoundWinner(gameBoard) || !hasEmptySpaces(gameBoard){
				break // end looping turns if someone wins or no more empty spaces
			}
			displayBoard(gameBoard)
			fmt.Println(P1.symbol, "score:", P1.score, P2.symbol, "symbol:",)
			applyMove(P1,P2,count, gameBoard)
			count ++
		}
		}
	switch{
	case P1.score > P2.score:
		winner = P1
	case P2.score > P1.score:
		winner = P2
	case P2.score == P1.score:
		winner = Player{"Draw! no one ",P1.score}
	}
	}


