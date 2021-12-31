// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/innovation-minning/message/npool"
	crud "github.com/NpoolPlatform/innovation-minning/pkg/crud/capital"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCapital(ctx context.Context, in *npool.CreateCapitalRequest) (*npool.CreateCapitalResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail create capital: %v", err)
		return &npool.CreateCapitalResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateCapital(ctx context.Context, in *npool.UpdateCapitalRequest) (*npool.UpdateCapitalResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail update capital: %v", err)
		return &npool.UpdateCapitalResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
