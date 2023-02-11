package cosmoschain_test

import (
	"testing"

	keepertest "github.com/cuttage/cosmoschain/testutil/keeper"
	"github.com/cuttage/cosmoschain/testutil/nullify"
	"github.com/cuttage/cosmoschain/x/cosmoschain"
	"github.com/cuttage/cosmoschain/x/cosmoschain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmoschainKeeper(t)
	cosmoschain.InitGenesis(ctx, *k, genesisState)
	got := cosmoschain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
