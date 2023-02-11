package keeper

import (
	"github.com/cuttage/cosmoschain/x/cosmoschain/types"
)

var _ types.QueryServer = Keeper{}
