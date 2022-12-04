package keeper

import (
	"context"

	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"mitoblockchaindev/x/mitoblockchaindev/types"
)

func (k msgServer) DeleteDiscountTokenStatus(goCtx context.Context, msg *types.MsgDeleteDiscountTokenStatus) (*types.MsgDeleteDiscountTokenStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	discountTokenStatus, exist := k.GetDiscountTokenStatus(ctx, msg.DiscountTokenStatusID)
	if !exist {
		return nil, sdkerrors.Wrapf(types.ErrID, "discount token status doesnt exist")
	}

	if msg.TokenID != discountTokenStatus.TokenID {
		return nil, sdkerrors.Wrapf(types.ErrID, "Disocunt token Id does not exist for which status with token Id %d was made", msg.TokenID)
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DiscountTokenStatusKey))
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, discountTokenStatus.Id)
	store.Delete(bz)

	return &types.MsgDeleteDiscountTokenStatusResponse{}, nil
}
