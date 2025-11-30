package main

// used: switch cases, for loops, "while" loops, explored different function implementations regarding data handling, conversion, computation.
import (
	"fmt"
	"math"
	// 	"math/rand"
		"strconv"
	//"regexp"
	"strings"
	"testing"
)

func testComputation (t *testing.T){
	//testing skeleton
}

type CustomError struct {
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error: %s", e.Message)
}

func generateNegativeInputFactorialError() *CustomError{
	return &CustomError {Message: "Negative number input, factorial not possible!"}
}
func applyComputation [T int | float64 ](x, y T, z string) T { //related to topic 1 and 2
	var result T
	switch z {
	case "x":
		result = x * y
	case "+":
		result = x + y
	case "-":
		result = x - y
	//case z == "%":
	//	result = x % y
	case "/":
		result = x / y
	}

	fmt.Println(x,z,y, "=",result)

	return result
}

func computeFactorial (x int) {
	if x >= 0{
	factorial := x
	for i := 1; i < x; i++ { //related to topic 1 and 2
		factorial *= x - i
	}
	fmt.Println("Factorial of",x,"=",factorial)	
	} else {
		fmt.Println(generateNegativeInputFactorialError())
	}
}

func sumUsingRange (nums []int) int {
	sum := 0
	for _ , n := range nums {
		sum += n
	}
	return sum
}


func isComposite (num int) bool{
	composite :=  false
	if num > 3 {
		numSqrt := int(math.Sqrt(float64(num)))  
		count := 0
		for i := 2; i <= numSqrt ; i++ {
			if num%i == 0{
				count ++
			}
		}
		if count > 0{
				composite = true
			}
	}

	return composite
}

func isPrime (num int) bool{
	prime := false
	if num <= 3{
		prime = true
	} else {
		numSqrt := int(math.Sqrt(float64(num)))  
		count := 0
		for i := 2; i <= numSqrt ; i++ {
			if num%i == 0{
				count ++
			}
					}
		if count == 0{
				prime = true
			}
	}
	return prime
}

func primesNComposite (nums []int ) ([]int, []int) {
	var primes, composites [ ] int
	for _, num := range nums{
		if isPrime(num){
			primes = append(primes, num)
		}
		if isComposite(num){
			composites = append(composites, num)
		}
	}
	return primes, composites 
	}

func identifyDatatypes (str string){
	strSlice := strings.Split(str, " ")
	for _ , char := range strSlice {
		if c, err := strconv.Atoi(char); err == nil{
			fmt.Printf("Char: %v  DataType: %T\n", char, c)
		} else if c, err := strconv.ParseFloat(char, 64); err == nil{
			fmt.Printf("Char: %v  DataType: %T\n", char, c)
		} else {
			fmt.Printf("Char: %v  DataType: %T\n", char, char)
		}
	}
}

func main() {
	fmt.Println()
	var f = 0
	numsArray := [9]int {2,3,4,5,6,7,8,9,10}
	applyComputation(4.5,8.9,"x")
	applyComputation(4,6,"-")

	fmt.Println("Enter Number to compute Factorial:")
	fmt.Scan(&f)
	computeFactorial(f)
	fmt.Println(sumUsingRange(numsArray[:]))
	p,c := primesNComposite(numsArray[:])
	fmt.Println("primes from num:", p, "\ncomposites from num:", c)
	testString := "60 B d 45.2"
	identifyDatatypes(testString)
}