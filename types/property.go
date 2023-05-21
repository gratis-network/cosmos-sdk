package types

import (
	"fmt"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

// NewProperty returns a new property vector with a set of balances.
func NewProperty(coins ...Coin) Property {
	return Property{
		Balances: coins,
	}
}

// String provides a human-readable representation of a PropertyVector
func (p Property) String() string {
	return fmt.Sprintf("%v", p.Balances)
}

func (p Property) ToAny() (*codectypes.Any, error) {
	return codectypes.NewAnyWithValue(&p)
}
