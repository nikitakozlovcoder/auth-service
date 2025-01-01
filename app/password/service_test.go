package password

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareSamePassword(t *testing.T) {
	// arrange
	passwordHash := "$2a$10$0b65BPYf9XRw0tYis0usQ.LhSg.546d0m/Cj7ux3YQTD5l5kDwOlS"
	password := "password"
	sut := NewPasswordComparer()

	//act
	isSame, err := sut.Compare(passwordHash, password)

	//assert
	assert.Empty(t, err)
	assert.True(t, isSame)
}

func TestCompareInvalidPassword(t *testing.T) {
	// arrange
	passwordHash := "$2a$10$0b65BPYf9XRw0tYis0usQ.LhSg.546d0m/Cj7ux3YQTD5l5kDwOlS"
	password := "123password123"
	sut := NewPasswordComparer()

	//act
	isSame, err := sut.Compare(passwordHash, password)

	//assert
	assert.Empty(t, err)
	assert.False(t, isSame)
}
