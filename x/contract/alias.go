package contract

import (
	"github.com/firmachain/FirmaChain/x/contract/internal/keeper"
	"github.com/firmachain/FirmaChain/x/contract/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewKeeper     = keeper.NewKeeper
	NewQuerier    = keeper.NewQuerier
	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec

	NewContract       = types.NewContract
	NewMsgAddContract = types.NewMsgAddContract
)

type (
	Keeper = keeper.Keeper

	Contract         = types.Contract
	MsgAddContract   = types.MsgAddContract
	QueryResContract = types.QueryResContract
)