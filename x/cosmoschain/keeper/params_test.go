package keeper_test

import (
	"testing"

	testkeeper "github.com/cuttage/cosmoschain/testutil/keeper"
	"github.com/cuttage/cosmoschain/x/cosmoschain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CosmoschainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
