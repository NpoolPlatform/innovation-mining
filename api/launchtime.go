// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/innovation-minning/message/npool"
	crud "github.com/NpoolPlatform/innovation-minning/pkg/crud/launchtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateLaunchTime(ctx context.Context, in *npool.CreateLaunchTimeRequest) (*npool.CreateLaunchTimeResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail create launchtime: %v", err)
		return &npool.CreateLaunchTimeResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateLaunchTime(ctx context.Context, in *npool.UpdateLaunchTimeRequest) (*npool.UpdateLaunchTimeResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail update launchtime: %v", err)
		return &npool.UpdateLaunchTimeResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
