// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/innovation-minning/message/npool"
	crud "github.com/NpoolPlatform/innovation-minning/pkg/crud/techniqueanalysis"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateTechniqueAnalysis(ctx context.Context, in *npool.CreateTechniqueAnalysisRequest) (*npool.CreateTechniqueAnalysisResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail create techniqueanalysis: %v", err)
		return &npool.CreateTechniqueAnalysisResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateTechniqueAnalysis(ctx context.Context, in *npool.UpdateTechniqueAnalysisRequest) (*npool.UpdateTechniqueAnalysisResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail update techniqueanalysis: %v", err)
		return &npool.UpdateTechniqueAnalysisResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
