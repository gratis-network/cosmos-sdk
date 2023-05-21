package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// module name
	ModuleName = "auth"

	// StoreKey is string representation of the store key for auth
	StoreKey = "acc"

	// FeeCollectorName the root string for the fee collector account address
	FeeCollectorName = "fee_collector"

	// QuerierRoute is the querier route for auth
	QuerierRoute = ModuleName

	// PropertyNftClassID is the NFT Class ID for user properties
	PropertyNftClassID = "gratis-property"
)

var (
	// AddressStoreKeyPrefix prefix for account-by-address store
	AddressStoreKeyPrefix = []byte{0x01}

	// GlobalAccountNumberKey param key for global account number
	GlobalAccountNumberKey = []byte("globalAccountNumber")

	// GlobalPropertyNumberKey param key for global property number
	GlobalPropertyNumberKey = []byte("globalPropertyNumber")

	// AccountNumberStoreKeyPrefix prefix for account-by-id store
	AccountNumberStoreKeyPrefix = []byte("accountNumber")
)

// AddressStoreKey turn an address to key used to get it from the account store
func AddressStoreKey(addr sdk.AccAddress) []byte {
	return append(AddressStoreKeyPrefix, addr.Bytes()...)
}

// AccountNumberStoreKey turn an account number to key used to get the account address from account store
func AccountNumberStoreKey(accountNumber uint64) []byte {
	return append(AccountNumberStoreKeyPrefix, sdk.Uint64ToBigEndian(accountNumber)...)
}
