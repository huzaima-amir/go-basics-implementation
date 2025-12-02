package main
// implementation of connect 3 game to apply topic 1, 2 and 3 concepts, including slices, structs, loops, string handling, pointers

// Players enter number of rounds -> players enter usernames -> outer loop for game round
//  (empty board prints out, with current scores, players get turn one by one and the round ends if one of them scores or the board gets filled up
// Winner is printed at the end of all the rounds using score comparison

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

type CustomError struct{
	Message string
} 

func (e *CustomError) Error() string {
    return fmt.Sprintf("Error: %s", e.Message)
}

func generatePositionOccupancyError() *CustomError {
	return &CustomError{ Message: "Position already occupied!",}
}

func applyMove(P1,P2 Player, count int,board [][] string) { // to apply game move to board
	var currentPlayer Player
	var c Coordinates

	if count%2 == 0{
		currentPlayer = P1
	} else {
		currentPlayer = P2
	}
	fmt.Println(currentPlayer.symbol, "make your move.\n x coordinate:")
	fmt.Scan(&c.x) // x coordinate for input symbol
	fmt.Println("y coordinate:")
	fmt.Scan(&c.y) // y coordinate for input symbol

	if board[c.y - 1][c.x - 1] == "_" {
		board[c.y - 1][c.x - 1] = currentPlayer.symbol
	} else {
		fmt.Println(generatePositionOccupancyError())
		applyMove(P1,P2,count,board)
	}
	checkRoundWinner(board, &P1, &P2)
}

// Helpers for ending round loop:
func checkRoundWinner(board [][]string, P1, P2 *Player) bool {
    lines := [][]string{}
	// verifying matches:
    // rows
    for i := 0; i < 3; i++ {
        lines = append(lines, board[i])
    }
    //columns
    for j := 0; j < 3; j++ {
        col := []string{board[0][j], board[1][j], board[2][j]}
        lines = append(lines, col)
    }
    //diagonals
    diag1 := []string{board[0][0], board[1][1], board[2][2]}
    diag2 := []string{board[0][2], board[1][1], board[2][0]}
    lines = append(lines, diag1, diag2)

    // check each line
    for _, line := range lines {
        if line[0] != "_" && line[0] == line[1] && line[1] == line[2] {
			switch line[0] {
				case P1.symbol:
					P1.score++
					fmt.Println("Round Winner:", P1.symbol)
				case P2.symbol:
					P2.score++
					fmt.Println("Round Winner:", P2.symbol)
			}
            return true
        }
    }
    return false
}


func hasEmptySpaces(board [][]string) bool { // to check for lack of empty spaces on board to end the round, if no one scores
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

func runConnect3() { // entire game
	var n int  // amount of game rounds
	P1 := Player{"", 0}
	P2 := Player{"", 0}
	var winner Player  // winner based on score comparison at the end of all rounds


	fmt.Println("Enter number of game rounds:")
	fmt.Scan(&n)
	fmt.Println("Enter Player 1 Symbol:")
	fmt.Scan(&P1.symbol)
	fmt.Println("Enter Player 2 Symbol:")
	fmt.Scan(&P2.symbol)

	for i := 0; i < n; i++ { //simulating game rounds
		gameBoard := createBoard()
		count := 0
		
		for { //loops forever unless if condition for break is true
			if checkRoundWinner(gameBoard, &P1, &P2) || !hasEmptySpaces(gameBoard){
				break // end looping turns if someone wins or no more empty spaces
			}
			displayBoard(gameBoard)
			fmt.Println(P1.symbol, "score:", P1.score, P2.symbol, "score:", P2.score)
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

	fmt.Println("Game Over!\n", winner.symbol, "won, score:", winner.score)
}

func main(){
	runConnect3()
	}
