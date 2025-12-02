package main

import (
	"fmt"
	"sync"
)

type CustomError struct{
	Message string
} 

func (e *CustomError) Error() string {
    return fmt.Sprintf("Error: %s", e.Message)
}
func generateLowBalanceError() *CustomError {
	return &CustomError{ Message: "Low Balance!"}
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
	mu sync.Mutex
}

func (b *BankAccount) MakeTransaction(spent float64){
	b.mu.Lock()
    defer b.mu.Unlock()

	b.balance -= spent
	b.CheckDebt()
}

func (b *BankAccount) WithdrawCash(withdrawn float64){

    b.mu.Lock()
    defer b.mu.Unlock()

	if b.balance >= withdrawn {
		b.balance -= withdrawn
	} else {
		fmt.Println(generateLowBalanceError()) //replace with error 
	}
}

func (b *BankAccount) Deposit(deposited float64){

	b.mu.Lock()
    defer b.mu.Unlock()

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
	b.mu.Lock()
    defer b.mu.Unlock()

	fmt.Println("Current Account Balance:", b.balance)
	return b.balance
}


// user events logging handled using channels:

func generateUserEvent(n int, c chan bool){
	var event string
	switch n{
	case 1:
		event = "User signed in"
	case 2:
		event = "User clicked button"
	case 3:
		event = "User signed out"
	}
	fmt.Println(event)
	c <- true
}

func main() {
	c := make(chan bool)
	
	go generateUserEvent(1, c)
	go generateUserEvent(3, c)
	go generateUserEvent(2, c)

	<- c
	<- c
	<- c
	fmt.Println("Events Logged.")

    account := &BankAccount{balance: 100}

    var wg sync.WaitGroup
    wg.Add(3)

    go func() {
        defer wg.Done()
        account.Deposit(50)
    }()

    go func() {
        defer wg.Done()
        account.MakeTransaction(30)
    }()

    go func() {
        defer wg.Done()
        account.WithdrawCash(20)
    }()

    wg.Wait()
    account.CheckBalance()

}