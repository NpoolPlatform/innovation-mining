package capital

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

func validateCapital(info *npool.Capital) error {
	if info.GetName() == "" {
		return xerrors.Errorf("invlaid name")
	}
	if info.GetLogo() == "" {
		return xerrors.Errorf("invlaid logo")
	}
	if info.GetIntroduction() == "" {
		return xerrors.Errorf("invlaid introduction")
	}
	return nil
}

func dbRowToCapital(row *ent.Capital) *npool.Capital {
	return &npool.Capital{
		ID:           row.ID.String(),
		Name:         row.Name,
		Logo:         row.Logo,
		Introduction: row.Introduction,
	}
}

func Create(ctx context.Context, in *npool.CreateCapitalRequest) (*npool.CreateCapitalResponse, error) {
	if err := validateCapital(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	info, err := cli.
		Capital.
		Create().
		SetName(in.GetInfo().GetName()).
		SetLogo(in.GetInfo().GetLogo()).
		SetIntroduction(in.GetInfo().GetIntroduction()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create capital: %v", err)
	}

	return &npool.CreateCapitalResponse{
		Info: dbRowToCapital(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateCapitalRequest) (*npool.UpdateCapitalResponse, error) {
	if err := validateCapital(in.GetInfo()); err != nil {
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
		Capital.
		UpdateOneID(id).
		SetName(in.GetInfo().GetName()).
		SetLogo(in.GetInfo().GetLogo()).
		SetIntroduction(in.GetInfo().GetIntroduction()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update capital: %v", err)
	}

	return &npool.UpdateCapitalResponse{
		Info: dbRowToCapital(info),
	}, nil
}
