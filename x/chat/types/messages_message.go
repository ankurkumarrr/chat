package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateMessage = "create_message"
	TypeMsgUpdateMessage = "update_message"
	TypeMsgDeleteMessage = "delete_message"
)

var _ sdk.Msg = &MsgCreateMessage{}

func NewMsgCreateMessage(creator string, body string) *MsgCreateMessage {
	return &MsgCreateMessage{
		Creator: creator,
		Body:    body,
	}
}

func (msg *MsgCreateMessage) Route() string {
	return RouterKey
}

func (msg *MsgCreateMessage) Type() string {
	return TypeMsgCreateMessage
}

func (msg *MsgCreateMessage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateMessage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateMessage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateMessage{}

func NewMsgUpdateMessage(creator string, id uint64, body string) *MsgUpdateMessage {
	return &MsgUpdateMessage{
		Id:      id,
		Creator: creator,
		Body:    body,
	}
}

func (msg *MsgUpdateMessage) Route() string {
	return RouterKey
}

func (msg *MsgUpdateMessage) Type() string {
	return TypeMsgUpdateMessage
}

func (msg *MsgUpdateMessage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateMessage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateMessage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteMessage{}

func NewMsgDeleteMessage(creator string, id uint64) *MsgDeleteMessage {
	return &MsgDeleteMessage{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteMessage) Route() string {
	return RouterKey
}

func (msg *MsgDeleteMessage) Type() string {
	return TypeMsgDeleteMessage
}

func (msg *MsgDeleteMessage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteMessage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteMessage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
