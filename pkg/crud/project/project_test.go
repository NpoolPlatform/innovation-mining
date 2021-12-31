package project

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

func assertProject(t *testing.T, actual, expected *npool.Project) {
	assert.Equal(t, actual.Name, expected.Name)
	assert.Equal(t, actual.GoodID, expected.GoodID)
	assert.Equal(t, actual.TeamID, expected.TeamID)
	assert.Equal(t, actual.CapitalIDs, expected.CapitalIDs)
	assert.Equal(t, actual.Introduction, expected.Introduction)
	assert.Equal(t, actual.Logo, expected.Logo)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	project := npool.Project{
		Name:         uuid.New().String(),
		GoodID:       uuid.New().String(),
		TeamID:       uuid.New().String(),
		CapitalIDs:   []string{uuid.New().String(), uuid.New().String()},
		Introduction: uuid.New().String(),
		Logo:         uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateProjectRequest{
		Info: &project,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertProject(t, resp.Info, &project)
	}

	project.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdateProjectRequest{
		Info: &project,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertProject(t, resp.Info, &project)
	}
}
