package keeper_test

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"testing"

	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
)

const (
	testClassID          = "kitty"
	testClassName        = "Crypto Kitty"
	testClassSymbol      = "kitty"
	testClassDescription = "Crypto Kitty"
	testClassURI         = "class uri"
	testClassURIHash     = "ae702cefd6b6a65fe2f991ad6d9969ed"
	testID               = "kitty1"
	testURI              = "kitty uri"
	testURIHash          = "229bfd3c1b431c14a526497873897108"
)

type TestSuite struct {
	suite.Suite

	app         *simapp.SimApp
	ctx         sdk.Context
	addrs       []sdk.AccAddress
	queryClient nft.QueryClient
}

func (s *TestSuite) SetupTest() {
	app := simapp.Setup(s.T(), false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	ctx = ctx.WithBlockHeader(tmproto.Header{Time: tmtime.Now()})
	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	nft.RegisterQueryServer(queryHelper, app.NFTKeeper)
	queryClient := nft.NewQueryClient(queryHelper)

	s.app = app
	s.ctx = ctx
	s.queryClient = queryClient
	s.addrs = simapp.AddTestAddrsIncremental(app, ctx, 3, sdk.NewInt(30000000))
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) TestSaveClass() {
	except := sdk.Class{
		Id:          testClassID,
		Name:        testClassName,
		Symbol:      testClassSymbol,
		Description: testClassDescription,
		Uri:         testClassURI,
		UriHash:     testClassURIHash,
	}
	err := s.app.NFTKeeper.SaveClass(s.ctx, except)
	s.Require().NoError(err)

	actual, has := s.app.NFTKeeper.GetClass(s.ctx, testClassID)
	s.Require().True(has)
	s.Require().EqualValues(except, actual)

	// a property NFT class is created by default
	classes := s.app.NFTKeeper.GetClasses(s.ctx)
	s.Require().Subset(classes, []*sdk.Class{&except})
}

func (s *TestSuite) TestUpdateClass() {
	class := sdk.Class{
		Id:          testClassID,
		Name:        testClassName,
		Symbol:      testClassSymbol,
		Description: testClassDescription,
		Uri:         testClassURI,
		UriHash:     testClassURIHash,
	}
	err := s.app.NFTKeeper.SaveClass(s.ctx, class)
	s.Require().NoError(err)

	noExistClass := sdk.Class{
		Id:          "kitty1",
		Name:        testClassName,
		Symbol:      testClassSymbol,
		Description: testClassDescription,
		Uri:         testClassURI,
		UriHash:     testClassURIHash,
	}

	err = s.app.NFTKeeper.UpdateClass(s.ctx, noExistClass)
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "nft class does not exist")

	except := sdk.Class{
		Id:          testClassID,
		Name:        "My crypto Kitty",
		Symbol:      testClassSymbol,
		Description: testClassDescription,
		Uri:         testClassURI,
		UriHash:     testClassURIHash,
	}

	err = s.app.NFTKeeper.UpdateClass(s.ctx, except)
	s.Require().NoError(err)

	actual, has := s.app.NFTKeeper.GetClass(s.ctx, testClassID)
	s.Require().True(has)
	s.Require().EqualValues(except, actual)
}

func (s *TestSuite) TestMint() {
	class := sdk.Class{
		Id:          testClassID,
		Name:        testClassName,
		Symbol:      testClassSymbol,
		Description: testClassDescription,
		Uri:         testClassURI,
		UriHash:     testClassURIHash,
	}
	err := s.app.NFTKeeper.SaveClass(s.ctx, class)
	s.Require().NoError(err)

	expNFT := sdk.NFT{
		ClassId: testClassID,
		Id:      testID,
		Uri:     testURI,
	}
	err = s.app.NFTKeeper.Mint(s.ctx, expNFT, s.addrs[0])
	s.Require().NoError(err)

	// test GetNFT
	actNFT, has := s.app.NFTKeeper.GetNFT(s.ctx, testClassID, testID)
	s.Require().True(has)
	s.Require().EqualValues(expNFT, actNFT)

	// test GetOwner
	owner := s.app.NFTKeeper.GetOwner(s.ctx, testClassID, testID)
	s.Require().True(s.addrs[0].Equals(owner))

	// test GetNFTsOfClass
	actNFTs := s.app.NFTKeeper.GetNFTsOfClass(s.ctx, testClassID)
	s.Require().EqualValues([]sdk.NFT{expNFT}, actNFTs)

	// test GetNFTsOfClassByOwner
	actNFTs = s.app.NFTKeeper.GetNFTsOfClassByOwner(s.ctx, testClassID, s.addrs[0])
	s.Require().EqualValues([]sdk.NFT{expNFT}, actNFTs)

	// test GetBalance
	balance := s.app.NFTKeeper.GetBalance(s.ctx, testClassID, s.addrs[0])
	s.Require().EqualValues(uint64(1), balance)

	// test GetTotalSupply
	supply := s.app.NFTKeeper.GetTotalSupply(s.ctx, testClassID)
	s.Require().EqualValues(uint64(1), supply)

	expNFT2 := sdk.NFT{
		ClassId: testClassID,
		Id:      testID + "2",
		Uri:     testURI + "2",
	}
	err = s.app.NFTKeeper.Mint(s.ctx, expNFT2, s.addrs[0])
	s.Require().NoError(err)

	// test GetNFTsOfClassByOwner
	actNFTs = s.app.NFTKeeper.GetNFTsOfClassByOwner(s.ctx, testClassID, s.addrs[0])
	s.Require().EqualValues([]sdk.NFT{expNFT, expNFT2}, actNFTs)

	// test GetBalance
	balance = s.app.NFTKeeper.GetBalance(s.ctx, testClassID, s.addrs[0])
	s.Require().EqualValues(uint64(2), balance)
}

func (s *TestSuite) TestBurn() {
	except := sdk.Class{
		Id:          testClassID,
		Name:        testClassName,
		Symbol:      testClassSymbol,
		Description: testClassDescription,
		Uri:         testClassURI,
		UriHash:     testClassURIHash,
	}
	err := s.app.NFTKeeper.SaveClass(s.ctx, except)
	s.Require().NoError(err)

	expNFT := sdk.NFT{
		ClassId: testClassID,
		Id:      testID,
		Uri:     testURI,
	}
	err = s.app.NFTKeeper.Mint(s.ctx, expNFT, s.addrs[0])
	s.Require().NoError(err)

	err = s.app.NFTKeeper.Burn(s.ctx, testClassID, testID)
	s.Require().NoError(err)

	// test GetNFT
	_, has := s.app.NFTKeeper.GetNFT(s.ctx, testClassID, testID)
	s.Require().False(has)

	// test GetOwner
	owner := s.app.NFTKeeper.GetOwner(s.ctx, testClassID, testID)
	s.Require().Nil(owner)

	// test GetNFTsOfClass
	actNFTs := s.app.NFTKeeper.GetNFTsOfClass(s.ctx, testClassID)
	s.Require().Empty(actNFTs)

	// test GetNFTsOfClassByOwner
	actNFTs = s.app.NFTKeeper.GetNFTsOfClassByOwner(s.ctx, testClassID, s.addrs[0])
	s.Require().Empty(actNFTs)

	// test GetBalance
	balance := s.app.NFTKeeper.GetBalance(s.ctx, testClassID, s.addrs[0])
	s.Require().EqualValues(uint64(0), balance)

	// test GetTotalSupply
	supply := s.app.NFTKeeper.GetTotalSupply(s.ctx, testClassID)
	s.Require().EqualValues(uint64(0), supply)
}

func (s *TestSuite) TestUpdate() {
	class := sdk.Class{
		Id:          testClassID,
		Name:        testClassName,
		Symbol:      testClassSymbol,
		Description: testClassDescription,
		Uri:         testClassURI,
		UriHash:     testClassURIHash,
	}
	err := s.app.NFTKeeper.SaveClass(s.ctx, class)
	s.Require().NoError(err)

	myNFT := sdk.NFT{
		ClassId: testClassID,
		Id:      testID,
		Uri:     testURI,
	}
	err = s.app.NFTKeeper.Mint(s.ctx, myNFT, s.addrs[0])
	s.Require().NoError(err)

	expNFT := sdk.NFT{
		ClassId: testClassID,
		Id:      testID,
		Uri:     "updated",
	}

	err = s.app.NFTKeeper.Update(s.ctx, expNFT)
	s.Require().NoError(err)

	// test GetNFT
	actNFT, has := s.app.NFTKeeper.GetNFT(s.ctx, testClassID, testID)
	s.Require().True(has)
	s.Require().EqualValues(expNFT, actNFT)
}

func (s *TestSuite) TestTransfer() {
	class := sdk.Class{
		Id:          testClassID,
		Name:        testClassName,
		Symbol:      testClassSymbol,
		Description: testClassDescription,
		Uri:         testClassURI,
		UriHash:     testClassURIHash,
	}
	err := s.app.NFTKeeper.SaveClass(s.ctx, class)
	s.Require().NoError(err)

	expNFT := sdk.NFT{
		ClassId: testClassID,
		Id:      testID,
		Uri:     testURI,
	}
	err = s.app.NFTKeeper.Mint(s.ctx, expNFT, s.addrs[0])
	s.Require().NoError(err)

	// valid owner
	err = s.app.NFTKeeper.Transfer(s.ctx, testClassID, testID, s.addrs[1])
	s.Require().NoError(err)

	// test GetOwner
	owner := s.app.NFTKeeper.GetOwner(s.ctx, testClassID, testID)
	s.Require().Equal(s.addrs[1], owner)

	balanceAddr0 := s.app.NFTKeeper.GetBalance(s.ctx, testClassID, s.addrs[0])
	s.Require().EqualValues(uint64(0), balanceAddr0)

	balanceAddr1 := s.app.NFTKeeper.GetBalance(s.ctx, testClassID, s.addrs[1])
	s.Require().EqualValues(uint64(1), balanceAddr1)

	// test GetNFTsOfClassByOwner
	actNFTs := s.app.NFTKeeper.GetNFTsOfClassByOwner(s.ctx, testClassID, s.addrs[1])
	s.Require().EqualValues([]sdk.NFT{expNFT}, actNFTs)
}

func (s *TestSuite) TestExportGenesis() {
	propertyClass := sdk.Class{
		Id: authtypes.PropertyNftClassID,
	}
	class := sdk.Class{
		Id:          testClassID,
		Name:        testClassName,
		Symbol:      testClassSymbol,
		Description: testClassDescription,
		Uri:         testClassURI,
		UriHash:     testClassURIHash,
	}
	err := s.app.NFTKeeper.SaveClass(s.ctx, class)
	s.Require().NoError(err)

	expNFT := sdk.NFT{
		ClassId: testClassID,
		Id:      testID,
		Uri:     testURI,
	}

	err = s.app.NFTKeeper.Mint(s.ctx, expNFT, s.addrs[0])
	s.Require().NoError(err)

	expGenesis := &nft.GenesisState{
		Classes: []*sdk.Class{&propertyClass, &class},
		Entries: []*nft.Entry{
			{
				Owner: s.addrs[0].String(),
				Nfts:  []*sdk.NFT{&expNFT},
			},
			{
				Owner: s.addrs[0].String(),
				Nfts:  []*sdk.NFT{&expNFT},
			},
		},
	}
	genesis := s.app.NFTKeeper.ExportGenesis(s.ctx)
	s.Require().Equal(expGenesis.Classes, genesis.Classes)
	// TODO: test property NFTs
	//json, err := s.app.AppCodec().MarshalJSON(genesis)
	//fmt.Printf("%s\n", json)
	//s.Require().Subset(genesis.Entries, expGenesis.Entries)
}

func (s *TestSuite) TestInitGenesis() {
	expClass := sdk.Class{
		Id:          testClassID,
		Name:        testClassName,
		Symbol:      testClassSymbol,
		Description: testClassDescription,
		Uri:         testClassURI,
		UriHash:     testClassURIHash,
	}
	expNFT := sdk.NFT{
		ClassId: testClassID,
		Id:      testID,
		Uri:     testURI,
	}
	expGenesis := &nft.GenesisState{
		Classes: []*sdk.Class{&expClass},
		Entries: []*nft.Entry{{
			Owner: s.addrs[0].String(),
			Nfts:  []*sdk.NFT{&expNFT},
		}},
	}
	s.app.NFTKeeper.InitGenesis(s.ctx, expGenesis)

	actual, has := s.app.NFTKeeper.GetClass(s.ctx, testClassID)
	s.Require().True(has)
	s.Require().EqualValues(expClass, actual)

	// test GetNFT
	actNFT, has := s.app.NFTKeeper.GetNFT(s.ctx, testClassID, testID)
	s.Require().True(has)
	s.Require().EqualValues(expNFT, actNFT)
}
