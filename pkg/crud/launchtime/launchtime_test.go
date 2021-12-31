package launchtime

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

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

func assertLaunchTime(t *testing.T, actual, expected *npool.LaunchTime) {
	assert.Equal(t, actual.NetworkName, expected.NetworkName)
	assert.Equal(t, actual.ProjectID, expected.ProjectID)
	assert.Equal(t, actual.NetworkVersion, expected.NetworkVersion)
	assert.Equal(t, actual.Introduction, expected.Introduction)
	assert.Equal(t, actual.LaunchTime, expected.LaunchTime)
	assert.Equal(t, actual.Incentive, expected.Incentive)
	assert.Equal(t, actual.IncentiveTotal, expected.IncentiveTotal)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	launchtime := npool.LaunchTime{
		ProjectID:      uuid.New().String(),
		NetworkName:    uuid.New().String(),
		NetworkVersion: uuid.New().String(),
		Introduction:   uuid.New().String(),
		LaunchTime:     uint32(time.Now().Unix()),
		Incentive:      true,
		IncentiveTotal: 10000000,
	}

	resp, err := Create(context.Background(), &npool.CreateLaunchTimeRequest{
		Info: &launchtime,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertLaunchTime(t, resp.Info, &launchtime)
	}

	launchtime.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdateLaunchTimeRequest{
		Info: &launchtime,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertLaunchTime(t, resp.Info, &launchtime)
	}
}
