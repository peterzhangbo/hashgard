package exchange

import (
	"github.com/hashgard/hashgard/x/exchange/keeper"
	"github.com/hashgard/hashgard/x/exchange/msgs"
	"github.com/hashgard/hashgard/x/exchange/queriers"
	"github.com/hashgard/hashgard/x/exchange/types"
)

type (
	Keeper = keeper.Keeper
	Order  = types.Order
	Orders = types.Orders
)

var (
	NewKeeper = keeper.NewKeeper

	RegisterCodec         = msgs.RegisterCodec
	NewMsgCreateOrder     = msgs.NewMsgCreateOrder
	NewMsgWithdrawalOrder = msgs.NewMsgWithdrawalOrder
	NewMsgTakeOrder       = msgs.NewMsgTakeOrder

	NewQueryOrderParams      = queriers.NewQueryOrderParams
	NewQueryOrdersParams     = queriers.NewQueryOrdersParams
	NewQueryFrozenFundParams = queriers.NewQueryFrozenFundParams
)

const (
	StoreKey          = types.StoreKey
	RouterKey         = types.RouterKey
	QuerierRoute      = types.QuerierRoute
	DefaultParamspace = types.DefaultParamspace
	DefaultCodespace  = types.DefaultCodespace
)
