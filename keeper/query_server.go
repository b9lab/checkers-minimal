package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/alice/checkers"
)

var _ checkers.QueryServer = queryServer{}

// NewQueryServerImpl returns an implementation of the module QueryServer.
func NewQueryServerImpl(k Keeper) checkers.QueryServer {
	return queryServer{k}
}

type queryServer struct {
	k Keeper
}

// GetGame defines the handler for the Query/GetGame RPC method.
func (qs queryServer) GetGame(ctx context.Context, req *checkers.QueryGetGameRequest) (*checkers.QueryGetGameResponse, error) {
	game, err := qs.k.StoredGames.Get(ctx, req.Index)
	if err == nil {
		return &checkers.QueryGetGameResponse{Game: &game}, nil
	}
	if errors.Is(err, collections.ErrNotFound) {
		return &checkers.QueryGetGameResponse{Game: nil}, nil
	}

	return nil, status.Error(codes.Internal, err.Error())
}
