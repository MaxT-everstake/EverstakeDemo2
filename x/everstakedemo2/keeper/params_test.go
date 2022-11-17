package keeper_test

import (
	"testing"

	testkeeper "Everstake_demo2/testutil/keeper"
	"Everstake_demo2/x/everstakedemo2/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.Everstakedemo2Keeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
