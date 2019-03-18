package internal

import (
	"fmt"
)

func CheckoutBook(s IStorageClient, isbn string, userID string) error {
	err := s.DecrementLibraryStock(isbn)
	if err != nil {
		return fmt.Errorf("Failed to decrement library stock", err)
	}

	err = s.IncrementUserStock(userID, isbn)
	if err != nil {
		return fmt.Errorf("Failed to increment user stock", err)
	}

	return nil
}

func ReturnBook(s IStorageClient, isbn string, userID string) error {
	err := s.IncrementLibraryStock(isbn)
	if err != nil {
		return fmt.Errorf("Failed to increment library stock", err)
	}

	err = s.DecrementUserStock(userID, isbn)
	if err != nil {
		return fmt.Errorf("Failed to decrement user stock", err)
	}

	return nil
}

func GetUserInventory(s IStorageClient, userID string) (Inventory, error) {
	return s.GetUserInventory(userID)
}
