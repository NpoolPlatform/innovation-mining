package capital

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

func assertCapital(t *testing.T, actual, expected *npool.Capital) {
	assert.Equal(t, actual.Name, expected.Name)
	assert.Equal(t, actual.Logo, expected.Logo)
	assert.Equal(t, actual.Introduction, expected.Introduction)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	capital := npool.Capital{
		Name:         uuid.New().String(),
		Logo:         uuid.New().String(),
		Introduction: uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateCapitalRequest{
		Info: &capital,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertCapital(t, resp.Info, &capital)
	}

	capital.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdateCapitalRequest{
		Info: &capital,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertCapital(t, resp.Info, &capital)
	}
}
