package types

// NewProperty returns a new property vector with a set of balances.
func NewProperty(coins ...Coin) Property {
	return Property{
		Balances: coins,
	}
}

// String provides a human-readable representation of a PropertyVector
func (v Property) String() string {
	return "property[" + v.Balances.String() + "]"
}
