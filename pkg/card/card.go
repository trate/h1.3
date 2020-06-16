// Package card works with the credit cards and transactions
package card

type Card struct {
	Id           int64
	Issuer       string
	Balance      int64
	Currency     string
	Number       string
	Transactions []Transaction
}

type Transaction struct {
	Id     int64
	Amount int64
	Date   int64
	MCC    string
	Status string
}

func AddTransaction(card *Card, transaction *Transaction) {
	card.Transactions = append(card.Transactions, *transaction)
}

func checkMCC(transaction Transaction, mcc []string) bool {
	for _, v := range mcc {
		if transaction.MCC == v {
			return true
		}
	}
	return false
}

func SumByMCC(transactions []Transaction, mcc []string) int64 {
	var result int64
	for _, v := range transactions {
		if checkMCC(v, mcc) {
			result += v.Amount
		}
	}
	return result
}

func LastNTransactions(card *Card, n int) []Transaction {
	transactions := make([]Transaction, 3, 6)
	copy(transactions, card.Transactions)
	if len(transactions) == 0 || n == 0 {
		transactions = nil
	} else if len(transactions) > n {
		transactions = transactions[len(transactions)-n:]
	}
	// reverse our slice
	for i := len(transactions)/2-1; i >= 0; i-- {
		opp := len(transactions)-1-i
		transactions[i], transactions[opp] = transactions[opp], transactions[i]
	}
	return transactions
}
