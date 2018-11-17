package test

import (
	"github.com/golang/mock/gomock"
	"github.com/hspearman/library-inventory/internal"
	"testing"
)

func Test_GetsInventory(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	s := NewMockIStorageClient(mockCtrl)

	s.EXPECT().GetInventory()

	_ = internal.GetInventory(s)
}
