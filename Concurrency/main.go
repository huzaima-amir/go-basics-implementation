package main

import (
	"fmt"
	"sync"
)



// -------- Bank interface code from the Interfaces folder, updated to implement mutexes so that the 
// balance cant be updated by multiple goroutines at once  -------------------------------

// custom errors:
type CustomError struct{
	Message string
} 

func (e *CustomError) Error() string {
    return fmt.Sprintf("Error: %s", e.Message)
}

func generateLowBalanceError() *CustomError {
	return &CustomError{ Message: "Low Balance!"}
}

func generateDebtError() *CustomError {
	return &CustomError{Message: "Account in debt! Transaction not possible."}
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
	mu sync.Mutex // to make sure balance is only accessed by 1 goroutine at a time
}

func (b *BankAccount) MakeTransaction(spent float64){
	b.mu.Lock()
    defer b.mu.Unlock()

	if !b.CheckDebt(){ // only works if balance is positive
		b.balance -= spent
		fmt.Println("Transaction of", spent,"made.")
	} else {
		fmt.Println(generateDebtError())
	}
	b.CheckDebt() // may cause debt
	b.CheckBalance()
}

func (b *BankAccount) WithdrawCash(withdrawn float64){

    b.mu.Lock()
    defer b.mu.Unlock()

	if b.balance >= withdrawn { // only works if balance is greater than or equal to withdrawal amount
		b.balance -= withdrawn
		fmt.Println("Amount of", withdrawn,"withdrawn.")

	} else {
		fmt.Println(generateLowBalanceError()) 
	}
	b.CheckDebt() //may cause debt
	b.CheckBalance()
}

func (b *BankAccount) Deposit(deposited float64){

	b.mu.Lock()
    defer b.mu.Unlock()

	b.balance += deposited // no extra conditions
	b.CheckDebt() // may diminish debt
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

	fmt.Println("Current Account Balance:", b.balance)  // to call within other functions (should lack mutex locks to prevent deadlocks)
	return b.balance
}

func (b *BankAccount) CheckBalanceIndependant() float64 { // to call separately (with mutex locks)
	b.mu.Lock()
    defer b.mu.Unlock()

	fmt.Println("Current Account Balance:", b.balance)
	return b.balance
}

// user events logging handled using channels and mutexes:


var ( 
	signedIn bool   // checking if signed in, shared state
	stateMu sync.Mutex  // mutex to protect events based 
	// on whether user is signed in or not.
)


func generateUserEvent(n int, c chan bool){
	var event string

	stateMu.Lock()
    defer stateMu.Unlock()

	switch n{
	case 1:
		if !signedIn {
			signedIn = true
			event = "User signed in"
		} else {
			event = "Sign in failed: already signed in" 
		}
		
	case 2:
		if signedIn {

			event = "User clicked button"
		} else {
			event = "Click failed: must be signed in"
		}
		
	case 3:
		if signedIn{
			signedIn = false
			event = "User signed out"
		} else {
			event = "Sign out failed: not signed in"
		}
		
	}
	fmt.Println(event) // "logging" the event after generation
	c <- true
}

func main() {
	c := make(chan bool)  
	
	go generateUserEvent(2, c)
	go generateUserEvent(3, c)
	go generateUserEvent(2, c)
	go generateUserEvent(1, c)

	<- c
	<- c
	<- c
	<- c

	fmt.Println("Events Logged.") // final message shouldnt print out unless all events are done.

    account := &BankAccount{balance: 100}

    var wg sync.WaitGroup
    wg.Add(3)

    go func() {
     	defer wg.Done()
		account.MakeTransaction(30)
    }()
   
	go func() {
     	defer wg.Done()
        account.Deposit(50)
    }()

	go func() {
    	defer wg.Done()
		account.WithdrawCash(10)
    }()

    wg.Wait()
    account.CheckBalance()

}
