package team

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

func validateTeam(info *npool.Team) error {
	if info.GetTeamName() == "" {
		return xerrors.Errorf("invlaid first name")
	}
	if _, err := uuid.Parse(info.GetLeaderID()); err != nil {
		return xerrors.Errorf("invalid leader id: %v", err)
	}
	for _, member := range info.GetMemberIDs() {
		if _, err := uuid.Parse(member); err != nil {
			return xerrors.Errorf("invalid member id: %v", err)
		}
	}
	return nil
}

func dbRowToTeam(row *ent.Team) *npool.Team {
	members := []string{}
	for _, member := range row.MemberIds {
		members = append(members, member.String())
	}

	return &npool.Team{
		ID:           row.ID.String(),
		TeamName:     row.TeamName,
		Logo:         row.Logo,
		LeaderID:     row.LeaderID.String(),
		MemberIDs:    members,
		Introduction: row.Introduction,
	}
}

func Create(ctx context.Context, in *npool.CreateTeamRequest) (*npool.CreateTeamResponse, error) {
	if err := validateTeam(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	members := []uuid.UUID{}
	for _, member := range in.GetInfo().GetMemberIDs() {
		members = append(members, uuid.MustParse(member))
	}

	info, err := cli.
		Team.
		Create().
		SetTeamName(in.GetInfo().GetTeamName()).
		SetLogo(in.GetInfo().GetLogo()).
		SetLeaderID(uuid.MustParse(in.GetInfo().GetLeaderID())).
		SetMemberIds(members).
		SetIntroduction(in.GetInfo().GetIntroduction()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create team: %v", err)
	}

	return &npool.CreateTeamResponse{
		Info: dbRowToTeam(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateTeamRequest) (*npool.UpdateTeamResponse, error) {
	if err := validateTeam(in.GetInfo()); err != nil {
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
		Team.
		UpdateOneID(id).
		SetTeamName(in.GetInfo().GetTeamName()).
		SetLogo(in.GetInfo().GetLogo()).
		SetIntroduction(in.GetInfo().GetIntroduction()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update team: %v", err)
	}

	return &npool.UpdateTeamResponse{
		Info: dbRowToTeam(info),
	}, nil
}
