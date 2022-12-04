package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateDiscountTokenStatus = "create_discount_token_status"

var _ sdk.Msg = &MsgCreateDiscountTokenStatus{}

func NewMsgCreateDiscountTokenStatus(creator string, tokenID uint64, timestamp string, status string) *MsgCreateDiscountTokenStatus {
	return &MsgCreateDiscountTokenStatus{
		Creator:   creator,
		TokenID:   tokenID,
		Timestamp: timestamp,
		Status:    status,
	}
}

func (msg *MsgCreateDiscountTokenStatus) Route() string {
	return RouterKey
}

func (msg *MsgCreateDiscountTokenStatus) Type() string {
	return TypeMsgCreateDiscountTokenStatus
}

func (msg *MsgCreateDiscountTokenStatus) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateDiscountTokenStatus) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateDiscountTokenStatus) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
