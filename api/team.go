// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/innovation-minning/message/npool"
	crud "github.com/NpoolPlatform/innovation-minning/pkg/crud/team"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateTeam(ctx context.Context, in *npool.CreateTeamRequest) (*npool.CreateTeamResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail create team: %v", err)
		return &npool.CreateTeamResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateTeam(ctx context.Context, in *npool.UpdateTeamRequest) (*npool.UpdateTeamResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail update team: %v", err)
		return &npool.UpdateTeamResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
