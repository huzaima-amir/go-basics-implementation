package main
// used: switch cases, for loops, "while" loops, explored different function implementations regarding data handling, conversion, computation.
import (
//	"errors"
	"fmt"
	//"math"
	// 	"math/rand"
//	"strconv"
	//"regexp"
)

func applyComputation (x, y float64, z string) { //related to topic 1 and 2
	result := 0.0
	switch {
	case z == "x":
		result = x * y
	case z == "+":
		result = x + y
	case z == "-":
		result = x - y
	//case z == "%": // check operations for float64
	//	result = x % y
	case z == "/":
		result = x / y
	}

	fmt.Println(x,z,y, "=",result)
}

func computeFactorial (x int) {
	if x >= 0{
	factorial := x
	for i := 1; i < x; i++ { //related to topic 1 and 2
		factorial *= x - i
	}
	fmt.Println("Factorial of",x,"=",factorial)	
	} else {
		fmt.Println("ERROR: Negative number entered, factorial not possible.") // replace with go errors import funcs
	}
}

//func parseExpression(){}

func main() {
	fmt.Println()
	var f, x, y = 0, 0.0, 0.0
	var z string
    fmt.Println("Enter Number 1:")//using scan for easier testing
	fmt.Scan(&x) //because scan in go writes to variable itself, instead of returning new value and assigning its value to the variable.
	fmt.Println("Enter Number 2:")
	fmt.Scan(&y)
	fmt.Println("Enter math operator:")
	fmt.Scan(&z)
	applyComputation(x,y,z)
	fmt.Println("Enter Number to compute Factorial:")
	fmt.Scan(&f)
	computeFactorial(f)
}