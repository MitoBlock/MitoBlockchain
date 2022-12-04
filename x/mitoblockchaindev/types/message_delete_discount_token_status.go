package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeleteDiscountTokenStatus = "delete_discount_token_status"

var _ sdk.Msg = &MsgDeleteDiscountTokenStatus{}

func NewMsgDeleteDiscountTokenStatus(creator string, discountTokenStatusID uint64, tokenID uint64) *MsgDeleteDiscountTokenStatus {
	return &MsgDeleteDiscountTokenStatus{
		Creator:               creator,
		DiscountTokenStatusID: discountTokenStatusID,
		TokenID:               tokenID,
	}
}

func (msg *MsgDeleteDiscountTokenStatus) Route() string {
	return RouterKey
}

func (msg *MsgDeleteDiscountTokenStatus) Type() string {
	return TypeMsgDeleteDiscountTokenStatus
}

func (msg *MsgDeleteDiscountTokenStatus) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteDiscountTokenStatus) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteDiscountTokenStatus) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
