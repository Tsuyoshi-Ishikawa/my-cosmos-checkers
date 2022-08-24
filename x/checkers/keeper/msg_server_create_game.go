package keeper

import (
	"context"
	"strconv"

	rules "github.com/alice/checkers/x/checkers/rules"
	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 新しいゲームのidを取得する
	nextGame, found := k.Keeper.GetNextGame(ctx)
	if !found {
		panic("NextGame not found")
	}
	newIndex := strconv.FormatUint(nextGame.IdValue, 10)

	// 新しいゲームデータを作成する
	newGame := rules.New()
	storedGame := types.StoredGame{
		Creator: msg.Creator,
		Index:   newIndex,
		Game:    newGame.String(),
		Turn:    rules.PieceStrings[newGame.Turn],
		Red:     msg.Red,
		Black:   msg.Black,
		MoveCount: 0,
		BeforeId:  types.NoFifoIdKey,
    AfterId:   types.NoFifoIdKey,
		Deadline: types.FormatDeadline(types.GetNextDeadline(ctx)),
		Winner:    rules.PieceStrings[rules.NO_PLAYER],
		Wager: msg.Wager,
	}

	// 新しいゲームデータの検証
	err := storedGame.Validate()
	if err != nil {
		return nil, err
	}

	// storedGameにbeforeGame、afterGameをセット
	k.Keeper.SendToFifoTail(ctx, &storedGame, &nextGame)
	// 新しいゲームデータの保存
	k.Keeper.SetStoredGame(ctx, storedGame)
	// この次のゲームのidを設定するためにincrementして保存する
	nextGame.IdValue++
	k.Keeper.SetNextGame(ctx, nextGame)

	// イベントを発火させる
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.StoredGameEventKey),
			sdk.NewAttribute(types.StoredGameEventCreator, msg.Creator),
			sdk.NewAttribute(types.StoredGameEventIndex, newIndex),
			sdk.NewAttribute(types.StoredGameEventRed, msg.Red),
			sdk.NewAttribute(types.StoredGameEventBlack, msg.Black),
			sdk.NewAttribute(types.StoredGameEventWager, strconv.FormatUint(msg.Wager, 10)),
		),
	)

	return &types.MsgCreateGameResponse{
		IdValue: newIndex,
	}, nil
}
