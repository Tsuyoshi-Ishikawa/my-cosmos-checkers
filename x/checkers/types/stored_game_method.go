package types

import (
	"errors"
	"fmt"

	"github.com/alice/checkers/x/checkers/rules"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// stored_game.pb.goは/protoのcompile結果なので直接いじるのは良くない
// なのでstored_gameにfunctionを設定したい場合はこのようにfileを作って対応してあげる。
// errorハンドリングはsdkerrorsで行い、エラーメッセージ(ErrInvalidCreatorなど)はx/checkers/types/errors.goに格納する
// 参考：https://github.com/cosmos/b9-checkers-academy-draft/blob/main/x/checkers/types/full_game.go

func (storedGame *StoredGame) GetCreatorAddress() (creator sdk.AccAddress, err error) {
	creator, errCreator := sdk.AccAddressFromBech32(storedGame.Creator)
	return creator, sdkerrors.Wrapf(errCreator, ErrInvalidCreator.Error(), storedGame.Creator)
}

func (storedGame *StoredGame) GetRedAddress() (red sdk.AccAddress, err error) {
	red, errRed := sdk.AccAddressFromBech32(storedGame.Red)
	return red, sdkerrors.Wrapf(errRed, ErrInvalidRed.Error(), storedGame.Red)
}

func (storedGame *StoredGame) GetBlackAddress() (black sdk.AccAddress, err error) {
	black, errBlack := sdk.AccAddressFromBech32(storedGame.Black)
	return black, sdkerrors.Wrapf(errBlack, ErrInvalidBlack.Error(), storedGame.Black)
}

func (storedGame *StoredGame) ParseGame() (game *rules.Game, err error) {
	game, errGame := rules.Parse(storedGame.Game)
	if errGame != nil {
		return nil, sdkerrors.Wrapf(errGame, ErrGameNotParseable.Error())
	}
	game.Turn = rules.StringPieces[storedGame.Turn].Player
	if game.Turn.Color == "" {
		return nil, sdkerrors.Wrapf(errors.New(fmt.Sprintf("Turn: %s", storedGame.Turn)), ErrGameNotParseable.Error())
	}
	return game, nil
}

func (storedGame StoredGame) Validate() (err error) {
	_, err = storedGame.GetCreatorAddress()
	if err != nil {
		return err
	}
	_, err = storedGame.ParseGame()
	if err != nil {
		return err
	}
	_, err = storedGame.GetRedAddress()
	if err != nil {
		return err
	}
	_, err = storedGame.GetBlackAddress()
	return err
}
