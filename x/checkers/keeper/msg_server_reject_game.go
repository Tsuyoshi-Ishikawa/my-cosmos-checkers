package keeper

import (
	"context"
	"strings"

	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	rules "github.com/alice/checkers/x/checkers/rules"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RejectGame(goCtx context.Context, msg *types.MsgRejectGame) (*types.MsgRejectGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	storedGame, found := k.Keeper.GetStoredGame(ctx, msg.IdValue)
	if !found {
			return nil, sdkerrors.Wrapf(types.ErrGameNotFound, "game not found %s", msg.IdValue)
	}

	// すでに勝敗がきまっていれば
	if storedGame.Winner != rules.PieceStrings[rules.NO_PLAYER] {
		return nil, types.ErrGameFinished
	}

	// 拒否しようとしているゲームが既に始まっているか or 関係ないプレーヤーによるゲーム拒否
	// の場合はエラーになる
	if strings.Compare(storedGame.Red, msg.Creator) == 0 {
    if 1 < storedGame.MoveCount { // Notice the use of the new field
        return nil, types.ErrRedAlreadyPlayed
    }
	} else if strings.Compare(storedGame.Black, msg.Creator) == 0 {
			if 0 < storedGame.MoveCount { // Notice the use of the new field
					return nil, types.ErrBlackAlreadyPlayed
			}
	} else {
			return nil, types.ErrCreatorNotPlayer
	}

	// Remove from the FIFO
	nextGame, found := k.Keeper.GetNextGame(ctx)
	if !found {
		panic("NextGame not found")
	}
	k.Keeper.RemoveFromFifo(ctx, &storedGame, &nextGame)

	k.Keeper.RemoveStoredGame(ctx, msg.IdValue)
	k.Keeper.SetNextGame(ctx, nextGame)

	ctx.EventManager().EmitEvent(
    sdk.NewEvent(sdk.EventTypeMessage,
        sdk.NewAttribute(sdk.AttributeKeyModule, "checkers"),
        sdk.NewAttribute(sdk.AttributeKeyAction, types.RejectGameEventKey),
        sdk.NewAttribute(types.RejectGameEventCreator, msg.Creator),
        sdk.NewAttribute(types.RejectGameEventIdValue, msg.IdValue),
    ),
)

	return &types.MsgRejectGameResponse{}, nil
}
