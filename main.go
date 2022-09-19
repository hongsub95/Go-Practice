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
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(word)

}
