package internal

func CheckoutBook(s IStorageClient, isbn string, userID string) {
	s.DecrementLibraryStock(isbn)
	s.IncrementUserStock(userID, isbn)
}

func ReturnBook(s IStorageClient, isbn string, userID string) {
	s.IncrementLibraryStock(isbn)
	s.DecrementUserStock(userID, isbn)
}

func GetUserInventory(s IStorageClient, userID string) Inventory {
	return s.GetUserInventory(userID)
}
