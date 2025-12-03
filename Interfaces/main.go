package main

import (
	"fmt"
	"math"
	//"time"
)

type CustomError struct{
	Message string
} 

func (e *CustomError) Error() string {
    return fmt.Sprintf("Error: %s", e.Message)
}
func generateLowBalanceError() *CustomError {
	return &CustomError{ Message: "Low Balance! Withdrawal not possible"}
}

func generateDebtError() *CustomError {
	return &CustomError{Message: "Account in debt! Transaction not possible."}
}

// shapes as an interface:
type Shapes interface {
	Area()  float64
	Perimeter()  float64
	ShapeDetails() string
}

type Rectangle struct {
	Length, Width float64 
}

type Circle struct {
	radius float64
}

func (r *Rectangle) Area () float64{
	return r.Length*r.Width
}

func (c *Circle) Area() float64 {
	return math.Pi*c.radius*c.radius
}

func (r *Rectangle)Perimeter() float64 {
	return 2*r.Length + 2*r.Width
}

func (c *Circle)Perimeter() float64 {
	return 2*c.radius*math.Pi
}

func (r *Rectangle)ShapeDetails(){
	fmt.Println("Rectangle with dimensions:", *r ,"has perimeter: ", r.Perimeter(), "and Area:", r.Area())
}

func (c *Circle)ShapeDetails(){
	fmt.Println("Circle with radius:", c.radius, ", has perimeter:", c.Perimeter(), "and Area:", c.Area())
}


// "Bank" as an interface:
type Bank interface {
	MakeTransaction() // minus money from acc, shouldnt work if there is debt on the acc, but can work if the amount paid is more than balance,
	//  which will put the user in debt

	WithdrawCash() //minus money in cash, should only work if withdrawal amount is less than or equal to balance
	Deposit() //add money
	CheckBalance() // returns current balance
	CheckDebt() // sets debt true if balance is negative and returns debt amount, and nil if no debt
}

type BankAccount struct {
	balance float64
	debt bool //should be set to true if balance is in the neg

}

func (b *BankAccount) MakeTransaction(spent float64){

	if !b.CheckDebt(){
		b.balance -= spent
		fmt.Println("Transaction of", spent,"made.")
	} else {
		fmt.Println(generateDebtError())
	}
	b.CheckDebt()
	b.CheckBalance()
}


func (b *BankAccount) WithdrawCash(withdrawn float64){

	if b.balance >= withdrawn {
		b.balance -= withdrawn
		fmt.Println("Amount of", withdrawn,"withdrawn.")

	} else {
		fmt.Println(generateLowBalanceError())
	}

	b.CheckDebt()
	b.CheckBalance()
}

func (b *BankAccount) Deposit(deposited float64){

	b.balance += deposited
	b.CheckDebt()
	fmt.Println("Amount of", deposited,"deposited.")
	b.CheckBalance()
}

func (b *BankAccount) CheckDebt() bool{
	b.debt = false
	if b.balance < 0{
		b.debt = true
	} 
	return b.debt
}

func (b *BankAccount) CheckBalance() float64 {

	fmt.Println("Current Account Balance:", b.balance)
	return b.balance
}

func main(){
	var c = Circle{2}
	c.ShapeDetails()
	var r = Rectangle{4,5.6}
	r.ShapeDetails()

	b := BankAccount{10.5,false}

	b.MakeTransaction(110)
	b.MakeTransaction(0.5)
	b.Deposit(100)
	b.WithdrawCash(5)
	b.Deposit(50)
	b.WithdrawCash(15)
	b.MakeTransaction(20)
	b.Deposit(150)
	
}