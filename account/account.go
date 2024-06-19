package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Account struct {
	Balance float64 `json:"balance"`
	Name    string  `json:"name"`
}

func New(name string) (a *Account, err error) {
	if name == "" {
		err := errors.New("account must have a name")

		return &Account{
			Name:    "default",
			Balance: 0.0,
		}, err
	}

	return &Account{
		Name:    name,
		Balance: 0.0,
	}, nil
}

func (a Account) DisplayData() {
	fmt.Printf("\n*~*~*~ Account Information ~*~*~*\n")
	fmt.Printf("Account Name: %v\n", a.Name)
	fmt.Printf("Account Balance: $%.2f\n\n", a.Balance)
}

func (a *Account) Withdraw(withdraw float64) (err error) {
	if withdraw > a.Balance || withdraw <= 0 {
		err := errors.New(
			"withdraw amount cannot be greater than balance or less than or equal to 0")

		return err
	}

	a.Balance -= withdraw

	fmt.Println("Withdraw successful!")
	fmt.Printf("Current balance: $%.2f\n", a.Balance)

	return nil
}

func (a *Account) Deposit(deposit float64) (err error) {
	if deposit <= 0 {
		err := errors.New("deposit amount cannot be less than or equal to 0")

		return err
	}

	a.Balance += deposit

	fmt.Printf("\nDeposit successful!\n")
	fmt.Printf("Current balance: $%.2f\n\n", a.Balance)

	return nil
}

func (a Account) Save() {
	fileName := a.Name + ".json"

	json, err := json.Marshal(a)

	if err != nil {
		fmt.Println(err)
	}

	os.WriteFile(fileName, json, 0644)
}
