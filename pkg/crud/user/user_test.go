package user

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

func assertUser(t *testing.T, actual, expected *npool.User) {
	assert.Equal(t, actual.FirstName, expected.FirstName)
	assert.Equal(t, actual.LastName, expected.LastName)
	assert.Equal(t, actual.Introduction, expected.Introduction)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	user := npool.User{
		FirstName:    uuid.New().String(),
		LastName:     uuid.New().String(),
		Introduction: uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateUserRequest{
		Info: &user,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertUser(t, resp.Info, &user)
	}

	user.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdateUserRequest{
		Info: &user,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertUser(t, resp.Info, &user)
	}
}
