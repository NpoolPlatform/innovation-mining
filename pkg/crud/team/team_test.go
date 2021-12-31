package team

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/innovation-minning/message/npool"
	"github.com/NpoolPlatform/innovation-minning/pkg/test-init" //nolint

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func assertTeam(t *testing.T, actual, expected *npool.Team) {
	assert.Equal(t, actual.TeamName, expected.TeamName)
	assert.Equal(t, actual.Logo, expected.Logo)
	assert.Equal(t, actual.LeaderID, expected.LeaderID)
	assert.Equal(t, actual.MemberIDs, expected.MemberIDs)
	assert.Equal(t, actual.Introduction, expected.Introduction)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	team := npool.Team{
		TeamName:     uuid.New().String(),
		Logo:         uuid.New().String(),
		LeaderID:     uuid.New().String(),
		MemberIDs:    []string{uuid.New().String(), uuid.New().String()},
		Introduction: uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateTeamRequest{
		Info: &team,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertTeam(t, resp.Info, &team)
	}

	team.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdateTeamRequest{
		Info: &team,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertTeam(t, resp.Info, &team)
	}
}
