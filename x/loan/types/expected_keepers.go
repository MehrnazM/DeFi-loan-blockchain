package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
	SendCoinsFromAccountToModule(ctx sdk.Context, addr sdk.AccAddress, moduleName string, collateral sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, moduleName string, addr sdk.AccAddress, collateral sdk.Coins) error
	SendCoins(ctx sdk.Context, lender sdk.AccAddress, borrower sdk.AccAddress, amount sdk.Coins) error
}
