package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"example.com/go-bank-structs/account"
)

func main() {
	displayBankInterface()
}

func displayBankInterface() {
	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What would you like to do today?")
		fmt.Println("1. Open a new account")
		fmt.Println("2. Bank with an existing account")
		fmt.Println("3. Exit")
		choice := getUserInput("Input selection here: ")

		switch choice {
		case 1:
			openNewAccount()
		case 2:
			bank()
		case 3:
			fmt.Println("Thank you for choosing Go Bank!")
			fmt.Println("Have a great day!")
			return
		default:
			fmt.Println("That was not an option, you dingus")
		}
	}
}

func openNewAccount() {
	var accountName string
	fmt.Print("What would you like to name your account?: ")
	fmt.Scanln(&accountName)

	account, err := account.New(accountName)

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		fmt.Println("Initializing account to default values...")
	}
	account.Save()
	fmt.Printf("\nNew account created!\n")

	account.DisplayData()

	fmt.Printf("Returning to main menu...\n\n")
}

func bank() {
	acc, err := getNote()

	if err != nil {
		fmt.Printf("\nERROR: %v\n\n", err)
		return
	}

	acc.DisplayData()
	displayBankOptions(acc)
}

func displayBankOptions(acc *account.Account) {
	fmt.Printf("What would you like to do with %v?:\n", acc.Name)
	fmt.Println("1. Deposit money")
	fmt.Println("2. Withdraw money")
	fmt.Println("3. Check account info")
	fmt.Println("4. Exit")
	choice := getUserInput("What would you like to do?: ")

	switch choice {
	case 1:
		var depositAmount float64
		fmt.Print("How much would you like to deposit?: ")
		fmt.Scan(&depositAmount)

		acc.Deposit(depositAmount)
	case 2:

	case 3:

	case 4:

	default:

		fmt.Println("This was not an option!")

	}
}

func getNote() (*account.Account, error) {
	var input string
	fmt.Print("Which account would you like to open?: ")
	fmt.Scan(&input)

	fileName := input + ".json"

	jsonFile, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	byteValue, _ := io.ReadAll(jsonFile)

	var acc account.Account

	json.Unmarshal(byteValue, &acc)

	return &acc, nil
}

func getUserInput(prompt string) int {
	var choice int
	fmt.Print(prompt)
	fmt.Scan(&choice)

	return choice
}
