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
}
