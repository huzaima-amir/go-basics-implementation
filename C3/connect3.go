package main
// implementation of connect 3 game to apply topic 1, 2 and 3 concepts, including slices, structs, loops

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

func applyMove(board [][] string) { // to apply game move to board
	
}

// Helpers for ending round loop:
func checkRoundWinner(board [][]string) bool { // check for a win to end round loop
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
	defer fmt.Println("Game Over!\n", "won with score:")
	fmt.Println("Enter number of game rounds:")
	fmt.Scan(&n)
	//fmt.Println(createBoard(n))
	fmt.Println("Enter Player 1 Symbol:")
	fmt.Scan(&P1.symbol)
	fmt.Println("Enter Player 2 Symbol:")
	fmt.Scan(&P2.symbol)

	for i := 0; i < n; i++ { //simulating game rounds
		gameBoard := createBoard()
		displayBoard(gameBoard)
		for {


			if checkRoundWinner(gameBoard) || hasEmptySpaces(gameBoard){
				break // end looping turns if someone wins or no more empty spaces
			}
		}
		}
	}


