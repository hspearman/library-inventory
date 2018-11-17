package test

import (
	"github.com/golang/mock/gomock"
	"testing"

	"github.com/hspearman/library-inventory/internal"
)

func Test_ChecksOutBook(t *testing.T) {
	userID := "1"
	isbn := "9780345404473"

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	s := NewMockIStorageClient(mockCtrl)

	s.EXPECT().DecrementLibraryStock(isbn)
	s.EXPECT().IncrementUserStock(userID, isbn)

	internal.CheckoutBook(s, isbn, userID)
}

func Test_ReturnsBook(t *testing.T) {
	userID := "1"
	isbn := "9780345404473"

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	s := NewMockIStorageClient(mockCtrl)

	s.EXPECT().IncrementLibraryStock(isbn)
	s.EXPECT().DecrementUserStock(userID, isbn)

	internal.ReturnBook(s, isbn, userID)
}
