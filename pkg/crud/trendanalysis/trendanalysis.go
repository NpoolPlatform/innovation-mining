package trendanalysis

import (
	"context"
	"time"

	"github.com/NpoolPlatform/innovation-minning/message/npool"

	"github.com/NpoolPlatform/innovation-minning/pkg/db"
	"github.com/NpoolPlatform/innovation-minning/pkg/db/ent"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

const (
	dbTimeout = 5 * time.Second
)

func validateTrendAnalysis(info *npool.TrendAnalysis) error {
	if info.GetTitle() == "" {
		return xerrors.Errorf("invlaid first name")
	}
	if info.GetContent() == "" {
		return xerrors.Errorf("invlaid last name")
	}
	if _, err := uuid.Parse(info.GetAuthorID()); err != nil {
		return xerrors.Errorf("invalid author id: %v", err)
	}
	if _, err := uuid.Parse(info.GetProjectID()); err != nil {
		return xerrors.Errorf("invalid project id: %v", err)
	}
	return nil
}

func dbRowToTrendAnalysis(row *ent.TrendAnalysis) *npool.TrendAnalysis {
	return &npool.TrendAnalysis{
		ID:        row.ID.String(),
		Title:     row.Title,
		Content:   row.Content,
		AuthorID:  row.AuthorID.String(),
		ProjectID: row.ProjectID.String(),
		Abstract:  row.Abstract,
	}
}

func Create(ctx context.Context, in *npool.CreateTrendAnalysisRequest) (*npool.CreateTrendAnalysisResponse, error) {
	if err := validateTrendAnalysis(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	info, err := cli.
		TrendAnalysis.
		Create().
		SetTitle(in.GetInfo().GetTitle()).
		SetContent(in.GetInfo().GetContent()).
		SetAbstract(in.GetInfo().GetAbstract()).
		SetAuthorID(uuid.MustParse(in.GetInfo().GetAuthorID())).
		SetProjectID(uuid.MustParse(in.GetInfo().GetProjectID())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create trendanalysis: %v", err)
	}

	return &npool.CreateTrendAnalysisResponse{
		Info: dbRowToTrendAnalysis(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateTrendAnalysisRequest) (*npool.UpdateTrendAnalysisResponse, error) {
	if err := validateTrendAnalysis(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	info, err := cli.
		TrendAnalysis.
		UpdateOneID(id).
		SetTitle(in.GetInfo().GetTitle()).
		SetContent(in.GetInfo().GetContent()).
		SetAbstract(in.GetInfo().GetAbstract()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update trendanalysis: %v", err)
	}

	return &npool.UpdateTrendAnalysisResponse{
		Info: dbRowToTrendAnalysis(info),
	}, nil
}
