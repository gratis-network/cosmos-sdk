package nft

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// ValidateGenesis check the given genesis state has no integrity issues
func ValidateGenesis(data GenesisState) error {
	for _, class := range data.Classes {
		if err := ValidateClassID(class.Id); err != nil {
			return err
		}
	}
	for _, entry := range data.Entries {
		for _, nft := range entry.Nfts {
			if err := ValidateNFTID(nft.Id); err != nil {
				return err
			}
			if _, err := sdk.AccAddressFromBech32(entry.Owner); err != nil {
				return err
			}
		}
	}
	return nil
}

// DefaultGenesisState - Return a default genesis state
func DefaultGenesisState() *GenesisState {
	classes := []*sdk.Class{{Id: types.PropertyNftClassID}}

	return &GenesisState{
		Classes: classes,
	}
}

// GetGenesisStateFromAppState returns x/nft GenesisState given raw application
// genesis state.
func GetGenesisStateFromAppState(cdc codec.JSONCodec, appState map[string]json.RawMessage) *GenesisState {
	var genesisState GenesisState

	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return &genesisState
}
