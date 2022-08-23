package keeper

import (
	"context"
	"strconv"

	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	rules "github.com/alice/checkers/x/checkers/rules"
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
	}

	// 新しいゲームデータの検証
	err := storedGame.Validate()
	if err != nil {
			return nil, err
	}

	// 新しいゲームデータの保存
	k.Keeper.SetStoredGame(ctx, storedGame)
	// この次のゲームのidを設定するためにincrementして保存する
	nextGame.IdValue++
	k.Keeper.SetNextGame(ctx, nextGame)

	return &types.MsgCreateGameResponse{
		IdValue: newIndex,
	}, nil
}
