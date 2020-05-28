package main

import (
	"fmt"

	"github.com/quavious/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account.Balance())
	account.Withdraw(20)
	//err := account.Withdraw(20)
	/*
		if err != nil {
			log.Fatalln(err)
		}*/
	fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)
}
