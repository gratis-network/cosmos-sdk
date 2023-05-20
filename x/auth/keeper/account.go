package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// NewAccountWithAddress implements AccountKeeperI.
func (ak AccountKeeper) NewAccountWithAddress(ctx sdk.Context, addr sdk.AccAddress) types.AccountI {
	acc := ak.proto()
	err := acc.SetAddress(addr)
	if err != nil {
		panic(err)
	}

	return ak.NewAccount(ctx, acc)
}

// NewAccount sets the next account number to a given account interface
func (ak AccountKeeper) NewAccount(ctx sdk.Context, acc types.AccountI) types.AccountI {
	if err := acc.SetAccountNumber(ak.GetNextAccountNumber(ctx)); err != nil {
		panic(err)
	}
	if ak.nftKeeper == nil {
		return acc
	}
	switch acc.(type) {
	case types.ModuleAccountI:
		return acc
	default:
		// create a default property
		nft, err := ak.newPropertyNFT(ctx, acc)
		if err != nil {
			return acc
		}
		// set the new property as primary
		acc.SetPropertyID(nft.Id)
		return acc
	}
}

// HasAccount implements AccountKeeperI.
func (ak AccountKeeper) HasAccount(ctx sdk.Context, addr sdk.AccAddress) bool {
	store := ctx.KVStore(ak.key)
	return store.Has(types.AddressStoreKey(addr))
}

// HasAccountAddressByID checks account address exists by id.
func (ak AccountKeeper) HasAccountAddressByID(ctx sdk.Context, id uint64) bool {
	store := ctx.KVStore(ak.key)
	return store.Has(types.AccountNumberStoreKey(id))
}

// GetAccount implements AccountKeeperI.
func (ak AccountKeeper) GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI {
	store := ctx.KVStore(ak.key)
	bz := store.Get(types.AddressStoreKey(addr))
	if bz == nil {
		return nil
	}

	return ak.decodeAccount(bz)
}

// GetAccountAddressById returns account address by id.
func (ak AccountKeeper) GetAccountAddressByID(ctx sdk.Context, id uint64) string {
	store := ctx.KVStore(ak.key)
	bz := store.Get(types.AccountNumberStoreKey(id))
	if bz == nil {
		return ""
	}
	return sdk.AccAddress(bz).String()
}

// GetAllAccounts returns all accounts in the accountKeeper.
func (ak AccountKeeper) GetAllAccounts(ctx sdk.Context) (accounts []types.AccountI) {
	ak.IterateAccounts(ctx, func(acc types.AccountI) (stop bool) {
		accounts = append(accounts, acc)
		return false
	})

	return accounts
}

// SetAccount implements AccountKeeperI.
func (ak AccountKeeper) SetAccount(ctx sdk.Context, acc types.AccountI) {
	addr := acc.GetAddress()
	store := ctx.KVStore(ak.key)

	// create a default property if not exist
	if len(acc.GetPropertyID()) == 0 {
		nft, err := ak.newPropertyNFT(ctx, acc)
		if err != nil {
			return
		}
		acc.SetPropertyID(nft.Id)
	}

	bz, err := ak.MarshalAccount(acc)
	if err != nil {
		panic(err)
	}

	store.Set(types.AddressStoreKey(addr), bz)
	store.Set(types.AccountNumberStoreKey(acc.GetAccountNumber()), addr.Bytes())
}

// RemoveAccount removes an account for the account mapper store.
// NOTE: this will cause supply invariant violation if called
func (ak AccountKeeper) RemoveAccount(ctx sdk.Context, acc types.AccountI) {
	addr := acc.GetAddress()
	store := ctx.KVStore(ak.key)
	store.Delete(types.AddressStoreKey(addr))
	store.Delete(types.AccountNumberStoreKey(acc.GetAccountNumber()))
}

// IterateAccounts iterates over all the stored accounts and performs a callback function.
// Stops iteration when callback returns true.
func (ak AccountKeeper) IterateAccounts(ctx sdk.Context, cb func(account types.AccountI) (stop bool)) {
	store := ctx.KVStore(ak.key)
	iterator := sdk.KVStorePrefixIterator(store, types.AddressStoreKeyPrefix)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		account := ak.decodeAccount(iterator.Value())

		if cb(account) {
			break
		}
	}
}

// newPropertyNFT creates a default property NFT
func (ak AccountKeeper) newPropertyNFT(ctx sdk.Context, acc types.AccountI) (sdk.NFT, error) {
	coins := sdk.NewCoins()
	property := sdk.NewProperty(coins...)
	return ak.MintProperty(ctx, acc, property)
}
