package main

import (
	"fmt"

	"example.com/go-bank-structs/account"
)

func main() {
	displayBankInterface()
	// account, err := account.New("test")

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// account.DisplayData()

	// err = account.Deposit(100)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// account.DisplayData()

	// err = account.Withdraw(50.24)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

}

func displayBankInterface() {
	fmt.Println("Welcome to Go Bank!")

	fmt.Println("What would you like to do today?")
	fmt.Println("1. Open a new account")
	fmt.Println("2. Bank with an existing account")
	fmt.Println("3. Exit")

	for {
		var choice int
		fmt.Print("Input selection here: ")
		fmt.Scan(&choice)

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

	fmt.Printf("\nNew account created!\n")

	account.DisplayData()
}

func bank() {
	fmt.Println("In bank function")
}
