package types
import (
  sdk "github.com/cosmos/cosmos-sdk/types"
  sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
  "github.com/google/uuid"
)
var _ sdk.Msg = &MsgCreatePost{}
type MsgCreatePost struct {
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Title   string         `json:"title" yaml:"title"`
  ID      string         `json:"id" yaml:"id"`
}

// NewMsgCreatePost creates the `MsgCreatePost` message
func NewMsgCreatePost(creator sdk.AccAddress, title string) MsgCreatePost {
  return MsgCreatePost{
    Creator: creator,
    Title:   title,
    ID:      uuid.New().String(),
  }
}

// Route ...
func (msg MsgCreatePost) Route() string {
  return RouterKey
}
// Type ...
func (msg MsgCreatePost) Type() string {
  return "CreatePost"
}
// GetSigners ...
func (msg MsgCreatePost) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}
// GetSignBytes ...
func (msg MsgCreatePost) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}
// ValidateBasic ...
func (msg MsgCreatePost) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}