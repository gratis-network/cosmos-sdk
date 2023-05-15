package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BankKeeper defines the contract needed for supply related APIs (noalias)
type BankKeeper interface {
	SendCoins(ctx sdk.Context, from, to sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}

type NftKeeper interface {
	HasClass(ctx sdk.Context, classID string) bool
	SaveClass(ctx sdk.Context, class sdk.Class) error
	Mint(ctx sdk.Context, token sdk.NFT, receiver sdk.AccAddress) error
	GetNFTsOfClassByOwner(ctx sdk.Context, classID string, owner sdk.AccAddress) (nfts []sdk.NFT)
	GetNFT(ctx sdk.Context, classID, nftID string) (sdk.NFT, bool)
	ParseData(token sdk.NFT, iface interface{}) (any, error)
	Update(ctx sdk.Context, token sdk.NFT) error
}
