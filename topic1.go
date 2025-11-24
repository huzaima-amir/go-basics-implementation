package main
import (
	"fmt"
//	"math"
// 	"math/rand"
//	"strconv"
//	"regexp"
)

// function for computing user input mathematical expression
func computeExpression(expression string, ) {
	parseExpression(expression) // parse string before computing
}
func parseExpression (expression string) {

}
func main() {
	fmt.Println()
	var name string
	var expression string
    fmt.Print("Enter your name: ")
    fmt.Scanln(&name) // because scanln in go writes to variable itself, instead of returning new value and assigning its value to the variable.
    fmt.Println("Hello,", name, "!\n Enter the expression you want to compute.")
	fmt.Scanln(&expression)
	computeExpression(expression)
}