package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeleteMembershipTokenStatus = "delete_membership_token_status"

var _ sdk.Msg = &MsgDeleteMembershipTokenStatus{}

func NewMsgDeleteMembershipTokenStatus(creator string, membershipTokenStatusID uint64, tokenID uint64) *MsgDeleteMembershipTokenStatus {
	return &MsgDeleteMembershipTokenStatus{
		Creator:                 creator,
		MembershipTokenStatusID: membershipTokenStatusID,
		TokenID:                 tokenID,
	}
}

func (msg *MsgDeleteMembershipTokenStatus) Route() string {
	return RouterKey
}

func (msg *MsgDeleteMembershipTokenStatus) Type() string {
	return TypeMsgDeleteMembershipTokenStatus
}

func (msg *MsgDeleteMembershipTokenStatus) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteMembershipTokenStatus) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteMembershipTokenStatus) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
