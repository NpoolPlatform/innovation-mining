package trendanalysis

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

func assertTrendAnalysis(t *testing.T, actual, expected *npool.TrendAnalysis) {
	assert.Equal(t, actual.Title, expected.Title)
	assert.Equal(t, actual.Content, expected.Content)
	assert.Equal(t, actual.AuthorID, expected.AuthorID)
	assert.Equal(t, actual.ProjectID, expected.ProjectID)
	assert.Equal(t, actual.Abstract, expected.Abstract)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	trendanalysis := npool.TrendAnalysis{
		Title:     uuid.New().String(),
		Content:   uuid.New().String(),
		AuthorID:  uuid.New().String(),
		ProjectID: uuid.New().String(),
		Abstract:  uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateTrendAnalysisRequest{
		Info: &trendanalysis,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertTrendAnalysis(t, resp.Info, &trendanalysis)
	}

	trendanalysis.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdateTrendAnalysisRequest{
		Info: &trendanalysis,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertTrendAnalysis(t, resp.Info, &trendanalysis)
	}
}
