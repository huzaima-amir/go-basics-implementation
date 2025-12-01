package main

import (
	"fmt"
	"math"
)


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
	MakeTransaction() //minus money from acc
	WithdrawCash() //minus money in cash, should only work if theres no debt.

	Deposit() //add money
	CheckBalance() // returns current balance
	CheckDebt() // sets debt true if balance is negative and returns debt amount, and nil if no debt
}

type BankAccount struct {
	balance float64
	debt bool //should be set to true if balance is in the neg
}

func (b *BankAccount) MakeTransaction(spent float64){
	b.balance -= spent
	b.CheckDebt()
}

func (b *BankAccount) WithdrawCash(withdrawn float64){
	if b.balance >= withdrawn {
		b.balance -= withdrawn
	} else {
		fmt.Println("Error! Low Balance")
	}
}

func (b *BankAccount) Deposit(deposited float64){
	b.balance += deposited
	b.CheckDebt()
}

func (b *BankAccount) CheckDebt() bool{
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
	var b = BankAccount{30, false}
	b.MakeTransaction(80)
	b.CheckBalance()
	b.WithdrawCash(40)
	fmt.Println(b.CheckDebt())

}