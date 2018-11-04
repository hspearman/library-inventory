package main

import "testing"

func Test_ChecksOutBook(t *testing.T) {
	var userID = "1"
	var isbn ISBN = "9780345404473"

	stock := getStock()
	amtStock := stock[isbn]
	if amtStock != DefaultStock {
		t.Errorf("Expected isbn to have %v in-stock, but got %v", DefaultStock, amtStock)
	}

	userStock := getUserStock(userID)
	amtStock = userStock[isbn]
	if amtStock != 0 {
		t.Errorf("Expected user to have 0 checked out for isbn, but got %v", amtStock)
	}

	checkoutBook(isbn, userID)

	stock = getStock()
	amtStock = stock[isbn]
	if amtStock != DefaultStock-1 {
		t.Errorf("Expected isbn to have %v in-stock, but got %v", DefaultStock-1, amtStock)
	}

	userStock = getUserStock(userID)
	amtStock = userStock[isbn]
	if amtStock != 1 {
		t.Errorf("Expected user to have 1 checked out for isbn, but got %v", amtStock)
	}
}

func Test_ReturnsBook(t *testing.T) {
	var userID = "1"
	var isbn ISBN = "9780345404473"

	stock := getStock()
	amtStock := stock[isbn]
	if amtStock != DefaultStock {
		t.Errorf("Expected isbn to have %v in-stock, but got %v", DefaultStock, amtStock)
	}

	userStock := getUserStock(userID)
	amtStock = userStock[isbn]
	if amtStock != 0 {
		t.Errorf("Expected user to have 0 checked out for isbn, but got %v", amtStock)
	}

	checkoutBook(isbn, userID)
	returnBook(isbn, userID)

	stock = getStock()
	amtStock = stock[isbn]
	if amtStock != DefaultStock {
		t.Errorf("Expected isbn to have %v in-stock, but got %v", DefaultStock, amtStock)
	}

	userStock = getUserStock(userID)
	amtStock = userStock[isbn]
	if amtStock != 0 {
		t.Errorf("Expected user to have 0 checked out for isbn, but got %v", amtStock)
	}
}
