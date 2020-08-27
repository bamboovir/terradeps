package basic

import (
	"github.com/bamboovir/terradeps/cmd/store/util"
	"github.com/bamboovir/terradeps/cmd/types"
)

// Store defination
type Store struct{}

// SetupRootLevel defination
func (s *Store) SetupRootLevel(state *types.State, prop *types.RootArgs) (*types.State, error) {
	util.SetupLogger(state.Logger, prop)
	state.RootArgs = prop
	return state, nil
}
