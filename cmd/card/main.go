package main

import (
	"fmt"
	"github.com/trate/h1.3/pkg/card"
)

func main() {
	t1 := &card.Transaction{
		Id:     1,
		Amount: -735_55,
		Date:   123456323,
		MCC:    "5411",
		Status: "InProgress",
	}
	t2 := &card.Transaction{
		Id:     2,
		Amount: 2_000_00,
		Date:   123423456323,
		MCC:    "0000",
		Status: "Done",
	}
	t3 := &card.Transaction{
		Id:     3,
		Amount: -1_203_91,
		Date:   123432223456323,
		MCC:    "5812",
		Status: "InProgress",
	}
	transactions := []card.Transaction{*t1, *t2}

	master := &card.Card{
		Id:           1,
		Issuer:       "MasterCard",
		Balance:      65_000,
		Currency:     "RUB",
		Number:       "5177827685644009",
		Transactions: transactions,
	}

	fmt.Println(master)

	fmt.Println("Add a new transaction.")

	card.AddTransaction(master, t3)
	fmt.Println(master)

	fmt.Println("Calculate the transactions amount for the restaurants and supermarkets")
	mcc := []string{"5411", "5812"}
	fmt.Println(float64(card.SumByMCC(master.Transactions, mcc)) / 100.0, "rubles")

	fmt.Println("Return last 3 transactions")
	fmt.Println(card.LastNTransactions(master,3))
	fmt.Println("Check that original struct is not changed")
	fmt.Println(master)
}
