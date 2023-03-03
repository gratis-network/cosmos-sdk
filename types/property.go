package types

import (
	"strings"
)

// NewProperty returns a new property vector with a set of balances.
func NewProperty(nft *NFT, coins ...Coin) Property {
	return Property{
		NFT:      nft,
		Balances: coins,
	}
}

// String provides a human-readable representation of a PropertyVector
func (v Property) String() string {
	// Build the string with a string builder
	var out strings.Builder
	out.WriteString(v.NFT.String())
	out.WriteString("balances:\n")
	out.WriteString(v.Balances.String())
	return out.String()
}
