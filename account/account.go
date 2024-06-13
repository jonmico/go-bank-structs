package account

import (
	"errors"
	"fmt"
)

type Account struct {
	balance float64
	name    string
}

func New(name string) (a *Account, err error) {
	if name == "" {
		err := errors.New("account must have a name")

		return &Account{
			name:    "default",
			balance: 0.0,
		}, err
	}

	return &Account{
		name:    name,
		balance: 0.0,
	}, nil
}

func (a Account) DisplayData() {
	fmt.Printf("Account Name: %v\n", a.name)
	fmt.Printf("Account Balance: $%.2f\n", a.balance)
}

func (a *Account) Withdraw(withdraw float64) (err error) {
	if withdraw > a.balance || withdraw <= 0 {
		err := errors.New(
			"withdraw amount cannot be greater than balance or less than or equal to 0")

		return err
	}

	a.balance -= withdraw

	fmt.Println("Withdraw successful!")
	fmt.Printf("Current balance: $%.2f\n", a.balance)

	return nil
}

func (a *Account) Deposit(deposit float64) (err error) {
	if deposit <= 0 {
		err := errors.New("deposit amount cannot be less than or equal to 0")

		return err
	}

	a.balance += deposit

	fmt.Println("Deposit successful!")
	fmt.Printf("Current balance: $%.2f\n", a.balance)

	return nil
}
