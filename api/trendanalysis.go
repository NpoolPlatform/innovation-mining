// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/innovation-minning/message/npool"
	crud "github.com/NpoolPlatform/innovation-minning/pkg/crud/trendanalysis"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateTrendAnalysis(ctx context.Context, in *npool.CreateTrendAnalysisRequest) (*npool.CreateTrendAnalysisResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail create trendanalysis: %v", err)
		return &npool.CreateTrendAnalysisResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateTrendAnalysis(ctx context.Context, in *npool.UpdateTrendAnalysisRequest) (*npool.UpdateTrendAnalysisResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail update trendanalysis: %v", err)
		return &npool.UpdateTrendAnalysisResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
