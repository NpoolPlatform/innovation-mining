// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/innovation-minning/message/npool"
	crud "github.com/NpoolPlatform/innovation-minning/pkg/crud/project"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateProject(ctx context.Context, in *npool.CreateProjectRequest) (*npool.CreateProjectResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail create project: %v", err)
		return &npool.CreateProjectResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateProject(ctx context.Context, in *npool.UpdateProjectRequest) (*npool.UpdateProjectResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail update project: %v", err)
		return &npool.UpdateProjectResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
