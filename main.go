package main

import (
	"fmt"

	"github.com/hongsub/learngo/Go-Practice/accounts"
)

func main() {
	account := accounts.NewAccount("hongsub")
	account.Deposit(10)
	fmt.Println(account.Balance())
}
