package keeper

import (
	"context"
	"fmt"
	"strings"

	rules "github.com/alice/checkers/x/checkers/rules"
	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ForfeitExpiredGames(goCtx context.Context) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	opponents := map[string]string{
		rules.PieceStrings[rules.BLACK_PLAYER]: rules.PieceStrings[rules.RED_PLAYER],
		rules.PieceStrings[rules.RED_PLAYER]:   rules.PieceStrings[rules.BLACK_PLAYER],
	}

	// Get FIFO information
	nextGame, found := k.GetNextGame(ctx)
	if !found {
		panic("NextGame not found")
	}

	storedGameId := nextGame.FifoHead
	var storedGame types.StoredGame
	for {
		// そもそも稼働中のゲームが一つもなければ終了
		if strings.Compare(storedGameId, types.NoFifoIdKey) == 0 {
			break
		}
		storedGame, found = k.GetStoredGame(ctx, storedGameId)
		if !found {
			panic("Fifo head game not found " + nextGame.FifoHead)
		}
		deadline, err := storedGame.GetDeadlineAsTime()
		if err != nil {
			panic(err)
		}
		if deadline.Before(ctx.BlockTime()) { // ゲームの有効期限がきれていれば、
			k.RemoveFromFifo(ctx, &storedGame, &nextGame)
			if storedGame.MoveCount <= 1 {
				// ゲームが作られただけでプレイされていなければ終了させる
				k.RemoveStoredGame(ctx, storedGameId)

				// すでに初手の人だけがplayしている場合は払い戻してあげる。
				if storedGame.MoveCount == 1 {
					k.MustRefundWager(ctx, &storedGame)
				}
			} else {
				// ゲームが途中ならば、勝者を決めてゲームを保存
				storedGame.Winner, found = opponents[storedGame.Turn]
				if !found {
					panic(fmt.Sprintf(types.ErrCannotFindWinnerByColor.Error(), storedGame.Turn))
				}
				k.MustPayWinnings(ctx, &storedGame) // 勝者に支払いをする
				k.SetStoredGame(ctx, storedGame) // SetStoredGameによってwinnerが決まったゲームはfifoから削除される
			}

			// イベント発火
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
					sdk.NewAttribute(sdk.AttributeKeyAction, types.ForfeitGameEventKey),
					sdk.NewAttribute(types.ForfeitGameEventIdValue, storedGameId),
					sdk.NewAttribute(types.ForfeitGameEventWinner, storedGame.Winner),
				),
			)

			// 次の稼働中のゲームをセットして同じ処理を行わせる
			storedGameId = nextGame.FifoHead
		} else {
			// 指定のゲームの有効期限がきれていないことがわかったので、
			// これ移行のゲームの有効期限が不要なので処理を終了させる
			break
		}
	}

	k.SetNextGame(ctx, nextGame)

}