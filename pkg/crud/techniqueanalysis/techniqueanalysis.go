package techniqueanalysis

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

func validateTechniqueAnalysis(info *npool.TechniqueAnalysis) error {
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

func dbRowToTechniqueAnalysis(row *ent.TechniqueAnalysis) *npool.TechniqueAnalysis {
	return &npool.TechniqueAnalysis{
		ID:        row.ID.String(),
		Title:     row.Title,
		Content:   row.Content,
		AuthorID:  row.AuthorID.String(),
		ProjectID: row.ProjectID.String(),
		Abstract:  row.Abstract,
	}
}

func Create(ctx context.Context, in *npool.CreateTechniqueAnalysisRequest) (*npool.CreateTechniqueAnalysisResponse, error) {
	if err := validateTechniqueAnalysis(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	info, err := cli.
		TechniqueAnalysis.
		Create().
		SetTitle(in.GetInfo().GetTitle()).
		SetContent(in.GetInfo().GetContent()).
		SetAbstract(in.GetInfo().GetAbstract()).
		SetAuthorID(uuid.MustParse(in.GetInfo().GetAuthorID())).
		SetProjectID(uuid.MustParse(in.GetInfo().GetProjectID())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create techniqueanalysis: %v", err)
	}

	return &npool.CreateTechniqueAnalysisResponse{
		Info: dbRowToTechniqueAnalysis(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateTechniqueAnalysisRequest) (*npool.UpdateTechniqueAnalysisResponse, error) {
	if err := validateTechniqueAnalysis(in.GetInfo()); err != nil {
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
		TechniqueAnalysis.
		UpdateOneID(id).
		SetTitle(in.GetInfo().GetTitle()).
		SetContent(in.GetInfo().GetContent()).
		SetAbstract(in.GetInfo().GetAbstract()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update techniqueanalysis: %v", err)
	}

	return &npool.UpdateTechniqueAnalysisResponse{
		Info: dbRowToTechniqueAnalysis(info),
	}, nil
}
