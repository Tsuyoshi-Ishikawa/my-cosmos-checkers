package keeper

import (
	"context"
	"strconv"
	"strings"

	rules "github.com/alice/checkers/x/checkers/rules"
	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) PlayMove(goCtx context.Context, msg *types.MsgPlayMove) (*types.MsgPlayMoveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// playしているgameのデータを取得
	storedGame, found := k.Keeper.GetStoredGame(ctx, msg.IdValue)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrGameNotFound, "game not found %s", msg.IdValue)
	}

	// 既に勝負がついているか？
	if storedGame.Winner != rules.PieceStrings[rules.NO_PLAYER] {
		return nil, types.ErrGameFinished
	}

	// 実行者がどのplayerかを判定してplayer変数に格納
	isRed := strings.Compare(storedGame.Red, msg.Creator) == 0
	isBlack := strings.Compare(storedGame.Black, msg.Creator) == 0
	var player rules.Player
	if !isRed && !isBlack {
		return nil, types.ErrCreatorNotPlayer
	} else if isRed && isBlack {
		player = rules.StringPieces[storedGame.Turn].Player
	} else if isRed {
		player = rules.RED_PLAYER
	} else {
		player = rules.BLACK_PLAYER
	}

	// 現在、実行者のターンなのか判定
	game, err := storedGame.ParseGame()
	if err != nil {
		panic(err.Error())
	}
	if !game.TurnIs(player) {
		return nil, types.ErrNotPlayerTurn
	}

	// Make the player pay the wager at the beginning
	err = k.Keeper.CollectWager(ctx, &storedGame)
	if err != nil {
		return nil, err
	}

	// Do it
	captured, moveErr := game.Move(
		rules.Pos{
			X: int(msg.FromX),
			Y: int(msg.FromY),
		},
		rules.Pos{
			X: int(msg.ToX),
			Y: int(msg.ToY),
		},
	)
	if moveErr != nil {
		return nil, sdkerrors.Wrapf(types.ErrWrongMove, moveErr.Error())
	}

	storedGame.MoveCount++
	storedGame.Deadline = types.FormatDeadline(types.GetNextDeadline(ctx))
	storedGame.Winner = rules.PieceStrings[game.Winner()] // まだ勝敗がついていなければrules.NO_PLAYERが入る

	// Remove from or send to the back of the FIFO
	nextGame, found := k.Keeper.GetNextGame(ctx)
	if !found {
		panic("NextGame not found")
	}
	if storedGame.Winner == rules.PieceStrings[rules.NO_PLAYER] {
		k.Keeper.SendToFifoTail(ctx, &storedGame, &nextGame)
	} else {
		k.Keeper.RemoveFromFifo(ctx, &storedGame, &nextGame)
		k.Keeper.MustPayWinnings(ctx, &storedGame) // Pay the winnings to the winner
	}

	// Save for the next play move
	storedGame.Game = game.String()
	storedGame.Turn = rules.PieceStrings[game.Turn]
	k.Keeper.SetStoredGame(ctx, storedGame)
	k.Keeper.SetNextGame(ctx, nextGame)

	ctx.GasMeter().ConsumeGas(types.PlayMoveGas, "Play a move")

	// イベントを発火させる
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.PlayMoveEventKey),
			sdk.NewAttribute(types.PlayMoveEventCreator, msg.Creator),
			sdk.NewAttribute(types.PlayMoveEventIdValue, msg.IdValue),
			sdk.NewAttribute(types.PlayMoveEventCapturedX, strconv.FormatInt(int64(captured.X), 10)),
			sdk.NewAttribute(types.PlayMoveEventCapturedY, strconv.FormatInt(int64(captured.Y), 10)),
			sdk.NewAttribute(types.PlayMoveEventWinner, rules.PieceStrings[game.Winner()]),
		),
	)

	// What to inform
	return &types.MsgPlayMoveResponse{
		IdValue:   msg.IdValue,
		CapturedX: int64(captured.X),
		CapturedY: int64(captured.Y),
		Winner:    rules.PieceStrings[game.Winner()],
	}, nil
}
