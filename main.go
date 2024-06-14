package main

import (
	"fmt"
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

	var choice int
	fmt.Print("Input selection here: ")
	fmt.Scan(&choice)
}
