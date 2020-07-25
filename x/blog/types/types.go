package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Post is a type containing Creator, Title, and ID
type Post struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Title   string         `json:"title" yaml:"title"`
	ID      string         `json:"id" yaml:"id"`
}
