package internal

// Inventory represents the library's stock of every book (ISBN mapped to # of books available)
type Inventory map[string]string

// DefaultStock represents the default stock for each book
var DefaultStock = "3"

var defaultInventory = Inventory{
	"9780345404473": DefaultStock, // Do Androids Dream of Electric Sheep?
	"9783641173081": DefaultStock, // Dune
	"9780679406419": DefaultStock, // The Complete Maus
}

func GetInventory(s IStorageClient) (Inventory, error) {
	return s.GetInventory()
}

func SetDefaultInventory(s IStorageClient) error {
	return s.SetInventory(defaultInventory)
}
