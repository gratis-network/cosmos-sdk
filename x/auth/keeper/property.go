package keeper

import (
	"fmt"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// GetProperty Fetch the primary property of the account
func (ak AccountKeeper) GetProperty(ctx sdk.Context, acc types.AccountI) (sdk.Property, error) {
	if acc == nil {
		return sdk.Property{}, fmt.Errorf("account does not exist")
	}
	nftKeeper := ak.GetNFTKeeper()
	propertyId := acc.GetPropertyID()
	nft, found := nftKeeper.GetNFT(ctx, types.PropertyNftClassID, propertyId)
	if !found {
		return sdk.Property{}, fmt.Errorf("property NFT with id %s does not exist", propertyId)
	}
	data, err := nftKeeper.ParseData(nft, &sdk.Property{})
	if err != nil {
		return sdk.Property{}, err
	}

	return *data.(*sdk.Property), nil
}

// UpdateProperty update the property of an account
func (ak AccountKeeper) UpdateProperty(ctx sdk.Context, acc types.AccountI, property sdk.Property) error {
	propertyId := acc.GetPropertyID()
	if len(propertyId) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrOutOfGas, "account %s does not have a property yet", acc.GetAddress())
	}
	nftKeeper := ak.GetNFTKeeper()
	nft, found := nftKeeper.GetNFT(ctx, types.PropertyNftClassID, propertyId)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "property NFT with id %s does not exist", propertyId)
	}

	newData, err := codectypes.NewAnyWithValue(&property)
	if err != nil {
		return err
	}

	newNFT := sdk.NFT{
		ClassId: nft.ClassId,
		Id:      nft.Id,
		Uri:     nft.Uri,
		UriHash: nft.UriHash,
		Data:    newData,
	}
	// save the new property
	err = nftKeeper.Update(ctx, newNFT)
	if err != nil {
		return err
	}
	return nil
}

// MintProperty mint a new property NFT
func (ak AccountKeeper) MintProperty(ctx sdk.Context, acc types.AccountI, property sdk.Property) (sdk.NFT, error) {
	if ak.nftKeeper == nil {
		return sdk.NFT{}, fmt.Errorf("AccountKeeper is null")
	}
	// check NFT class
	if !ak.nftKeeper.HasClass(ctx, types.PropertyNftClassID) {
		// add property NFT class (only once)
		ak.nftKeeper.SaveClass(ctx, sdk.Class{Id: types.PropertyNftClassID})
	}
	data, err := codectypes.NewAnyWithValue(&property)
	if err != nil {
		return sdk.NFT{}, fmt.Errorf("fail to encode property, error is %v", err)
	}
	id := fmt.Sprintf("%d", acc.GetAccountNumber())
	nft := sdk.NFT{
		ClassId: types.PropertyNftClassID,
		Id:      id,
		Uri:     "gratis/property/" + id,
		Data:    data,
	}
	err = ak.nftKeeper.Mint(ctx, nft, acc.GetAddress())
	if err != nil {
		return sdk.NFT{}, fmt.Errorf("fail to mint NFT, error is %v", err)
	}

	return nft, nil
}

// AddBalanceToProperty add given amount to the balances of the primary property of an account
func (ak AccountKeeper) AddBalanceToProperty(ctx sdk.Context, acc types.AccountI, amt sdk.Coins) error {
	if !amt.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, amt.String())
	}
	property, err := ak.GetProperty(ctx, acc)
	if err != nil {
		return err
	}
	balances := property.Balances
	newBalance := balances.Add(amt...)
	property.Balances = newBalance

	return ak.UpdateProperty(ctx, acc, property)
}
