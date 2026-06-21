package main

// To avoid compiler errors, the import is commented out until you write your code.
// Once you start writing your Withdraw function, uncomment the line below:
// import "errors"

type BankAccount struct {
	Owner   string
	Balance float64
}

// TODO: Implement the Deposit method (increases balance).
func (b *BankAccount) Deposit(amount float64) {
	// Write your code here
}

// TODO: Implement the Withdraw method. 
// If amount is greater than the balance, return an error (use errors.New("insufficient funds")).
// Hint: You will need to uncomment the import statement or use another package.
func (b *BankAccount) Withdraw(amount float64) error {
	// Write your code here
	return nil
}
