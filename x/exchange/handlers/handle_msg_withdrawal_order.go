package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/exchange/keeper"
	"github.com/hashgard/hashgard/x/exchange/msgs"
)

func HandleMsgWithdrawalOrder(ctx sdk.Context, k keeper.Keeper, msg msgs.MsgWithdrawalOrder) sdk.Result {
	// 做状态性判断

	// 执行状态变更

	// 返回 tags


	return sdk.Result{}
}