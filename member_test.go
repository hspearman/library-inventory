package libraryinventory

import "testing"

func Test_ChecksOutBook(t *testing.T) {
	var isbn ISBN = "9780345404473"

	stock := getStock()

	amtStock := stock[isbn]
	if amtStock != DefaultStock {
		t.Errorf("Expected isbn to have %v in-stock, but got %v", DefaultStock, amtStock)
	}

	checkoutBook(isbn)

	stock = getStock()

	amtStock = stock[isbn]
	if amtStock != DefaultStock-1 {
		t.Errorf("Expected isbn to have %v in-stock, but got %v", DefaultStock-1, amtStock)
	}
}

func Test_ReturnsBook(t *testing.T) {
	var isbn ISBN = "9780345404473"

	stock := getStock()

	amtStock := stock[isbn]
	if amtStock != DefaultStock {
		t.Errorf("Expected isbn to have %v in-stock, but got %v", DefaultStock, amtStock)
	}

	returnBook(isbn)

	stock = getStock()

	amtStock = stock[isbn]
	if amtStock != DefaultStock+1 {
		t.Errorf("Expected isbn to have %v in-stock, but got %v", DefaultStock+1, amtStock)
	}
}
