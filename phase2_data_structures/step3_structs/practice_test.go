package main

import "testing"

func TestBankAccount(t *testing.T) {
	acc := BankAccount{Owner: "Shubham", Balance: 100.0}

	acc.Deposit(50.0)
	if acc.Balance != 150.0 {
		t.Errorf("Expected balance 150.0, got %f", acc.Balance)
	}

	err := acc.Withdraw(30.0)
	if err != nil {
		t.Errorf("Unexpected error during withdrawal: %v", err)
	}
	if acc.Balance != 120.0 {
		t.Errorf("Expected balance 120.0, got %f", acc.Balance)
	}

	err = acc.Withdraw(200.0)
	if err == nil {
		t.Error("Expected error 'insufficient funds', got nil")
	}
}
