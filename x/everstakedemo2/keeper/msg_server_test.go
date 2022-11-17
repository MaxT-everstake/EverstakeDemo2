package keeper_test

import (
	"context"
	"testing"

	keepertest "Everstake_demo2/testutil/keeper"
	"Everstake_demo2/x/everstakedemo2/keeper"
	"Everstake_demo2/x/everstakedemo2/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.Everstakedemo2Keeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
