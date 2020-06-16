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

func TranslateMCC(code string) string {
	// представим, что mcc читается из файла (научимся позже)
	mcc := map[string]string{
		"5411": "Супермаркеты",
		"5812": "Рестораны",
	}

	if value, ok := mcc[code]; !ok {
		return "Категория не указана"
	} else {
		return value
	}
}
