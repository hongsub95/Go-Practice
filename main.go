package main

import (
	"fmt"

	"github.com/hongsub/learngo/Go-Practice/mydict"
)

// account:나의 계좌 balance:잔고 deposit:입금 withdraw:출금
/*
func main() {
	account := accounts.NewAccount("hongsub")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}
*/

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	err := dictionary.Add("third", "third word")
	if err != nil {
		fmt.Println(err)
	}
	def, err2 := dictionary.Search("third")
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(def)
	}
}
