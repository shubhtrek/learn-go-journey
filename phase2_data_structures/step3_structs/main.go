package main

import "fmt"

type Account struct {
	Owner   string
	Balance float64
}

// Value receiver - cannot change the balance
func (a Account) GetBalance() float64 {
	return a.Balance
}

// Pointer receiver - can modify the balance
func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}

func main() {
	acc := Account{Owner: "Shubham", Balance: 500}
	fmt.Printf("%s's account balance: %.2f\n", acc.Owner, acc.GetBalance())

	acc.Deposit(150.50)
	fmt.Printf("New balance after deposit: %.2f\n", acc.GetBalance())

	// 🎯 MINI CHALLENGE:
	// Add a `Withdraw` method to the Account struct.
	// Make sure it returns an error/message if the account balance goes below zero!
}
