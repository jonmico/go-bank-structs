package main

import (
	"fmt"

	"example.com/go-bank-structs/account"
)

func main() {
	account, err := account.New("test")

	if err != nil {
		fmt.Println(err)
		return
	}

	account.DisplayData()

	err = account.Deposit(100)

	if err != nil {
		fmt.Println(err)
		return
	}

	account.DisplayData()

	err = account.Withdraw(50.24)

	if err != nil {
		fmt.Println(err)
		return
	}

}
