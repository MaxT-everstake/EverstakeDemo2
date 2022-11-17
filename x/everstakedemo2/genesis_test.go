package everstakedemo2_test

import (
	"testing"

	keepertest "Everstake_demo2/testutil/keeper"
	"Everstake_demo2/testutil/nullify"
	"Everstake_demo2/x/everstakedemo2"
	"Everstake_demo2/x/everstakedemo2/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.Everstakedemo2Keeper(t)
	everstakedemo2.InitGenesis(ctx, *k, genesisState)
	got := everstakedemo2.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
