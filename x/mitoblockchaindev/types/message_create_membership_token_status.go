package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateMembershipTokenStatus = "create_membership_token_status"

var _ sdk.Msg = &MsgCreateMembershipTokenStatus{}

func NewMsgCreateMembershipTokenStatus(creator string, tokenID uint64, timestamp string, status string) *MsgCreateMembershipTokenStatus {
	return &MsgCreateMembershipTokenStatus{
		Creator:   creator,
		TokenID:   tokenID,
		Timestamp: timestamp,
		Status:    status,
	}
}

func (msg *MsgCreateMembershipTokenStatus) Route() string {
	return RouterKey
}

func (msg *MsgCreateMembershipTokenStatus) Type() string {
	return TypeMsgCreateMembershipTokenStatus
}

func (msg *MsgCreateMembershipTokenStatus) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateMembershipTokenStatus) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateMembershipTokenStatus) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
