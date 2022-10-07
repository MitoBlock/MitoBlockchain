package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateDiscountToken = "create_discount_token"

var _ sdk.Msg = &MsgCreateDiscountToken{}

func NewMsgCreateDiscountToken(creator string, timestamp string, activityName string, score string, message string, discountValue string, eligibleCompanies string, itemType string, expiryDate string) *MsgCreateDiscountToken {
  return &MsgCreateDiscountToken{
		Creator: creator,
    Timestamp: timestamp,
    ActivityName: activityName,
    Score: score,
    Message: message,
    DiscountValue: discountValue,
    EligibleCompanies: eligibleCompanies,
    ItemType: itemType,
    ExpiryDate: expiryDate,
	}
}

func (msg *MsgCreateDiscountToken) Route() string {
  return RouterKey
}

func (msg *MsgCreateDiscountToken) Type() string {
  return TypeMsgCreateDiscountToken
}

func (msg *MsgCreateDiscountToken) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgCreateDiscountToken) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateDiscountToken) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}

