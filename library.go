package libraryinventory

type stock map[ISBN]int

// DefaultStock represents the default stock for each book
var DefaultStock = 3

var libraryStock = stock{
	"9780345404473": DefaultStock, // Do Androids Dream of Electric Sheep?
	"9783641173081": DefaultStock, // Dune
	"9780679406419": DefaultStock, // The Complete Maus
}

func getStock() stock {
	return libraryStock
}
