package keeper

import (
	"github.com/mitoblock/mitoblockchain/x/mitoblockchain/types"
)

var _ types.QueryServer = Keeper{}
