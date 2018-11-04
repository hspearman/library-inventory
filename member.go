package libraryinventory

// ISBN represents an international standard book number
type ISBN string

func checkoutBook(isbn ISBN) {
	libraryStock[isbn] = libraryStock[isbn] - 1
}

func returnBook(isbn ISBN) {
	libraryStock[isbn] = libraryStock[isbn] + 1
}
