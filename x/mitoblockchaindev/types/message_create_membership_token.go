package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateMembershipToken = "create_membership_token"

var _ sdk.Msg = &MsgCreateMembershipToken{}

func NewMsgCreateMembershipToken(creator string, timestamp string, activityName string, score string, message string, membershipDuration string, eligibleCompanies string, expiryDate string) *MsgCreateMembershipToken {
	return &MsgCreateMembershipToken{
		Creator:            creator,
		Timestamp:          timestamp,
		ActivityName:       activityName,
		Score:              score,
		Message:            message,
		MembershipDuration: membershipDuration,
		EligibleCompanies:  eligibleCompanies,
		ExpiryDate:         expiryDate,
	}
}

func (msg *MsgCreateMembershipToken) Route() string {
	return RouterKey
}

func (msg *MsgCreateMembershipToken) Type() string {
	return TypeMsgCreateMembershipToken
}

func (msg *MsgCreateMembershipToken) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateMembershipToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateMembershipToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
