package store

import (
	"github.com/bamboovir/terradeps/cmd/store/basic"
	"github.com/bamboovir/terradeps/cmd/types"
)

// Store defination
type Store interface {
	// SetupRootLevel defination
	SetupRootLevel(state *types.State, prop *types.RootArgs) (*types.State, error)
}

// NewState defination
func NewState() *types.State {
	return &types.State{}
}

// NewBasicStore defination
func NewBasicStore() Store {
	return &basic.Store{}
}
