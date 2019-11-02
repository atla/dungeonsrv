package game

import (
	"github.com/atla/lotd/users"
)

// Character ... default entity for characters
type Character struct {
	Object
	Player* Player
}

// NewAvatar ... creates and returns a new room instance
func NewAvatar() *Avatar {
	return &Avatar{
		LastKnownRoom: nil,
	}
}
