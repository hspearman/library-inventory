package libraryinventory

import "testing"

func Test_GetsStock(t *testing.T) {
	stock := getStock()

	if stock == nil {
		t.Errorf("Expected stock but got nil")
	}

	isbns := []ISBN{
		"9780345404473",
		"9783641173081",
		"9780679406419",
	}

	for _, isbn := range isbns {
		amtStock := stock[isbn]
		if amtStock != DefaultStock {
			t.Errorf("Expected isbn to have %v in-stock, but got %v", DefaultStock, amtStock)
		}
	}
}
