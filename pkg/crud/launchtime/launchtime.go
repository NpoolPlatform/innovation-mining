package launchtime

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

func validateLaunchTime(info *npool.LaunchTime) error {
	if _, err := uuid.Parse(info.GetProjectID()); err != nil {
		return xerrors.Errorf("invalid project id: %v", err)
	}
	if info.GetNetworkName() == "" {
		return xerrors.Errorf("invlaid first name")
	}
	if info.GetNetworkVersion() == "" {
		return xerrors.Errorf("invlaid last name")
	}
	return nil
}

func dbRowToLaunchTime(row *ent.LaunchTime) *npool.LaunchTime {
	return &npool.LaunchTime{
		ID:             row.ID.String(),
		ProjectID:      row.ProjectID.String(),
		NetworkName:    row.NetworkName,
		NetworkVersion: row.NetworkVersion,
		Introduction:   row.Introduction,
		LaunchTime:     row.LaunchTime,
		Incentive:      row.Incentive,
		IncentiveTotal: row.IncentiveTotal,
	}
}

func Create(ctx context.Context, in *npool.CreateLaunchTimeRequest) (*npool.CreateLaunchTimeResponse, error) {
	if err := validateLaunchTime(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	info, err := cli.
		LaunchTime.
		Create().
		SetProjectID(uuid.MustParse(in.GetInfo().GetProjectID())).
		SetNetworkName(in.GetInfo().GetNetworkName()).
		SetNetworkVersion(in.GetInfo().GetNetworkVersion()).
		SetIntroduction(in.GetInfo().GetIntroduction()).
		SetLaunchTime(in.GetInfo().GetLaunchTime()).
		SetIncentive(in.GetInfo().GetIncentive()).
		SetIncentiveTotal(in.GetInfo().GetIncentiveTotal()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create launchtime: %v", err)
	}

	return &npool.CreateLaunchTimeResponse{
		Info: dbRowToLaunchTime(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateLaunchTimeRequest) (*npool.UpdateLaunchTimeResponse, error) {
	if err := validateLaunchTime(in.GetInfo()); err != nil {
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
		LaunchTime.
		UpdateOneID(id).
		SetNetworkName(in.GetInfo().GetNetworkName()).
		SetNetworkVersion(in.GetInfo().GetNetworkVersion()).
		SetIntroduction(in.GetInfo().GetIntroduction()).
		SetLaunchTime(in.GetInfo().GetLaunchTime()).
		SetIncentiveTotal(in.GetInfo().GetIncentiveTotal()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update launchtime: %v", err)
	}

	return &npool.UpdateLaunchTimeResponse{
		Info: dbRowToLaunchTime(info),
	}, nil
}
