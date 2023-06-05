package mockauthdatastore

import (
	"github.com/golang/mock/gomock"
)

type MockBCryptHasher struct {
	ctrl *gomock.Controller
}

func NewMockBCryptHasher(ctrl *gomock.Controller) *MockBCryptHasher {
	return &MockBCryptHasher{ctrl: ctrl}
}

func (m *MockBCryptHasher) HashPassword(password string) ([]byte, error) {
	hashedPassword, _ := make([]uint8, 0), gomock.Nil()
	return hashedPassword, nil
}
