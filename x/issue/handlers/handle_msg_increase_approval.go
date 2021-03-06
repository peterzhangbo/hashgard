package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/issue/utils"

	"github.com/hashgard/hashgard/x/issue/keeper"
	"github.com/hashgard/hashgard/x/issue/msgs"
)

//Handle MsgIssueIncreaseApproval
func HandleMsgIssueIncreaseApproval(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgIssueIncreaseApproval) sdk.Result {

	if err := keeper.IncreaseApproval(ctx, msg.Sender, msg.Spender, msg.IssueId, msg.Amount); err != nil {
		return err.Result()
	}

	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
		Tags: utils.GetIssueTags(msg.IssueId, msg.Sender),
	}
}
