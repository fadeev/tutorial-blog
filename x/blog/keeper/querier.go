package keeper

import (
	// this line is used by starport scaffolding
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/fadeev/blog/x/blog/types"
)

// NewQuerier creates a new querier for blog clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		// this line is used by starport scaffolding # 2
		case types.QueryListPosts:
			return listPosts(ctx, k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown blog query endpoint")
		}
	}
}
func listPosts(ctx sdk.Context, k Keeper) ([]byte, error) {
	var postList []types.Post
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.PostPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var post types.Post
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &post)
		postList = append(postList, post)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, postList)
	return res, nil
}
