package cache

import (
	"errors"

	types "github.com/prysmaticlabs/prysm/consensus-types/primitives"
)

// ErrNotCommittee will be returned when a cache object is not a pointer to
// a Committee struct.
var ErrNotCommittee = errors.New("object is not a committee struct")

// Committees defines the shuffled committees seed.
type Committees struct {
	CommitteeCount  uint64
	Seed            [32]byte
	ShuffledIndices []types.ValidatorIndex
	SortedIndices   []types.ValidatorIndex
}
