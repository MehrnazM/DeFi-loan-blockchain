package keeper

import (
	"context"

	"loan/x/loan/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RequestLoan(goCtx context.Context, msg *types.MsgRequestLoan) (*types.MsgRequestLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	loan := types.Loan{
		Amount:     msg.Amount,
		Collateral: msg.Collateral,
		Fee:        msg.Fee,
		Deadline:   msg.Deadline,
		Borrower:   msg.Creator,
		State:      "requested",
	}

	borrower, _ := sdk.AccAddressFromBech32(msg.Creator)

	// Get the collateral as sdk.Coins
	collateral, err := sdk.ParseCoinsNormalized(loan.Collateral)
	if err != nil {
		panic(err)
	}

	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, borrower, types.ModuleName, collateral)
	if sdkError != nil {
		return nil, sdkError
	}

	k.AppendLoan(ctx, loan)

	return &types.MsgRequestLoanResponse{}, nil
}