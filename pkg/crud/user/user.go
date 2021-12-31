package user

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

func validateUser(info *npool.User) error {
	if info.GetFirstName() == "" {
		return xerrors.Errorf("invlaid first name")
	}
	if info.GetLastName() == "" {
		return xerrors.Errorf("invlaid last name")
	}
	return nil
}

func dbRowToUser(row *ent.User) *npool.User {
	return &npool.User{
		ID:           row.ID.String(),
		FirstName:    row.FirstName,
		LastName:     row.LastName,
		Introduction: row.Introduction,
	}
}

func Create(ctx context.Context, in *npool.CreateUserRequest) (*npool.CreateUserResponse, error) {
	if err := validateUser(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	info, err := cli.
		User.
		Create().
		SetFirstName(in.GetInfo().GetFirstName()).
		SetLastName(in.GetInfo().GetLastName()).
		SetIntroduction(in.GetInfo().GetIntroduction()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create user: %v", err)
	}

	return &npool.CreateUserResponse{
		Info: dbRowToUser(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateUserRequest) (*npool.UpdateUserResponse, error) {
	if err := validateUser(in.GetInfo()); err != nil {
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
		User.
		UpdateOneID(id).
		SetFirstName(in.GetInfo().GetFirstName()).
		SetLastName(in.GetInfo().GetLastName()).
		SetIntroduction(in.GetInfo().GetIntroduction()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update user: %v", err)
	}

	return &npool.UpdateUserResponse{
		Info: dbRowToUser(info),
	}, nil
}
