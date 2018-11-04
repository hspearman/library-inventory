package main

var stockByUser = map[string]Stock{}

func checkoutBook(isbn ISBN, userID string) {
	allStock[isbn] = allStock[isbn] - 1

	if _, exists := stockByUser[userID]; !exists {
		stockByUser[userID] = Stock{}
	}

	if _, exists := stockByUser[userID][isbn]; exists {
		stockByUser[userID][isbn]++
	} else {
		stockByUser[userID][isbn] = 1
	}
}

func returnBook(isbn ISBN, userID string) {
	allStock[isbn] = allStock[isbn] + 1

	stockByUser[userID][isbn]--
}

func getUserStock(userID string) Stock {
	return stockByUser[userID]
}
