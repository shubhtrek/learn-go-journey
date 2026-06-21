package main

import "errors"

type BankAccount struct {
	Owner   string
	Balance float64
}

// TODO: Implement the Deposit method (increases balance).
func (b *BankAccount) Deposit(amount float64) {
	// Write your code here
}

// TODO: Implement the Withdraw method. 
// If amount is greater than the balance, return an error: errors.New("insufficient funds")
func (b *BankAccount) Withdraw(amount float64) error {
	// Write your code here
	return nil
}
