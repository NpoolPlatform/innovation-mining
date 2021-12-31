package project

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

func validateProject(info *npool.Project) error {
	if info.GetName() == "" {
		return xerrors.Errorf("invlaid first name")
	}
	if _, err := uuid.Parse(info.GetGoodID()); err != nil {
		return xerrors.Errorf("invalid good id: %v", err)
	}
	if _, err := uuid.Parse(info.GetTeamID()); err != nil {
		return xerrors.Errorf("invalid team id: %v", err)
	}
	for _, capital := range info.GetCapitalIDs() {
		if _, err := uuid.Parse(capital); err != nil {
			return xerrors.Errorf("invalid capital id: %v", err)
		}
	}
	if info.GetIntroduction() == "" {
		return xerrors.Errorf("invalid introduction")
	}
	if info.GetLogo() == "" {
		return xerrors.Errorf("invalid logo")
	}
	return nil
}

func dbRowToProject(row *ent.Project) *npool.Project {
	capitals := []string{}
	for _, capital := range row.CapitalIds {
		capitals = append(capitals, capital.String())
	}

	return &npool.Project{
		ID:           row.ID.String(),
		Name:         row.Name,
		GoodID:       row.GoodID.String(),
		TeamID:       row.TeamID.String(),
		CapitalIDs:   capitals,
		Introduction: row.Introduction,
		Logo:         row.Logo,
	}
}

func Create(ctx context.Context, in *npool.CreateProjectRequest) (*npool.CreateProjectResponse, error) {
	if err := validateProject(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	capitals := []uuid.UUID{}
	for _, capital := range in.GetInfo().GetCapitalIDs() {
		capitals = append(capitals, uuid.MustParse(capital))
	}

	info, err := cli.
		Project.
		Create().
		SetName(in.GetInfo().GetName()).
		SetGoodID(uuid.MustParse(in.GetInfo().GetGoodID())).
		SetTeamID(uuid.MustParse(in.GetInfo().GetTeamID())).
		SetCapitalIds(capitals).
		SetIntroduction(in.GetInfo().GetIntroduction()).
		SetLogo(in.GetInfo().GetLogo()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create project: %v", err)
	}

	return &npool.CreateProjectResponse{
		Info: dbRowToProject(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateProjectRequest) (*npool.UpdateProjectResponse, error) {
	if err := validateProject(in.GetInfo()); err != nil {
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

	capitals := []uuid.UUID{}
	for _, capital := range in.GetInfo().GetCapitalIDs() {
		capitals = append(capitals, uuid.MustParse(capital))
	}

	info, err := cli.
		Project.
		UpdateOneID(id).
		SetCapitalIds(capitals).
		SetIntroduction(in.GetInfo().GetIntroduction()).
		SetLogo(in.GetInfo().GetLogo()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update project: %v", err)
	}

	return &npool.UpdateProjectResponse{
		Info: dbRowToProject(info),
	}, nil
}
