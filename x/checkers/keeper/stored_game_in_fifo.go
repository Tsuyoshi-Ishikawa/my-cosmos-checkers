package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/alice/checkers/x/checkers/types"
)

// WARN It does not save game or info.
// gameのsaveを行うのはあくまでこの関数を呼び出した側で行う
func (k Keeper) RemoveFromFifo(ctx sdk.Context, game *types.StoredGame, info *types.NextGame) {
	// このゲームの前にゲームが存在するならば
	if game.BeforeId != types.NoFifoIdKey {
		beforeElement, found := k.GetStoredGame(ctx, game.BeforeId)
		if !found {
			panic("Element before in Fifo was not found")
		}
		beforeElement.AfterId = game.AfterId
		k.SetStoredGame(ctx, beforeElement)
		if game.AfterId == types.NoFifoIdKey {
			info.FifoTail = beforeElement.Index
		}
	} else if info.FifoHead == game.Index { // このゲームがFIFOの先頭ならば
		info.FifoHead = game.AfterId
	}
	// このゲームの後ろにゲームがなければ
	if game.AfterId != types.NoFifoIdKey {
		afterElement, found := k.GetStoredGame(ctx, game.AfterId)
		if !found {
			panic("Element after in Fifo was not found")
		}
		afterElement.BeforeId = game.BeforeId
		k.SetStoredGame(ctx, afterElement)
		if game.BeforeId == types.NoFifoIdKey { // そもそもFIFOにこのゲームしか登録されていない場合
			info.FifoHead = afterElement.Index
		}
		// このゲームが連なるFIFOの一番最後ならば
	} else if info.FifoTail == game.Index {
		info.FifoTail = game.BeforeId
	}
	game.BeforeId = types.NoFifoIdKey
	game.AfterId = types.NoFifoIdKey
}

// WARN It does not save game or info.
// gameのsaveを行うのはあくまでこの関数を呼び出した側で行う
func (k Keeper) SendToFifoTail(ctx sdk.Context, game *types.StoredGame, info *types.NextGame) {
	// まだゲームがFIFOに一つも設定されていなければ
	if info.FifoHead == types.NoFifoIdKey && info.FifoTail == types.NoFifoIdKey {
		game.BeforeId = types.NoFifoIdKey
		game.AfterId = types.NoFifoIdKey
		info.FifoHead = game.Index
		info.FifoTail = game.Index
	} else if info.FifoHead == types.NoFifoIdKey || info.FifoTail == types.NoFifoIdKey {
		panic("Fifo should have both head and tail or none")
	} else if info.FifoTail == game.Index { // 既にこのゲームがFIFOの一番最後ならば
		// Nothing to do, already at tail
	} else {
		// もしこのゲームがFIFOの一番最後より前にあるならば
		// FIFOにあるこのゲームを一旦削除して、FIFOの一番最後に追加する
		
		// Snip game out
		k.RemoveFromFifo(ctx, game, info)

		// Now add to tail
		currentTail, found := k.GetStoredGame(ctx, info.FifoTail)
		if !found {
			panic("Current Fifo tail was not found")
		}
		currentTail.AfterId = game.Index
		k.SetStoredGame(ctx, currentTail)

		game.BeforeId = currentTail.Index
		info.FifoTail = game.Index
	}
}