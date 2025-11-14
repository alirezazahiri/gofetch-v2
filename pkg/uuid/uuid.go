package uuid

import (
	"github.com/google/uuid"
)

func New() string {
	return uuid.New().String()
}

func NewV4() string {
	return uuid.New().String()
}
