package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/cuttage/cosmoschain/testutil/keeper"
	"github.com/cuttage/cosmoschain/x/cosmoschain/keeper"
	"github.com/cuttage/cosmoschain/x/cosmoschain/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CosmoschainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
