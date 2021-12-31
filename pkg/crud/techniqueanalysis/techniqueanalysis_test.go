package techniqueanalysis

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

func assertTechniqueAnalysis(t *testing.T, actual, expected *npool.TechniqueAnalysis) {
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

	techniqueanalysis := npool.TechniqueAnalysis{
		Title:     uuid.New().String(),
		Content:   uuid.New().String(),
		AuthorID:  uuid.New().String(),
		ProjectID: uuid.New().String(),
		Abstract:  uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateTechniqueAnalysisRequest{
		Info: &techniqueanalysis,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertTechniqueAnalysis(t, resp.Info, &techniqueanalysis)
	}

	techniqueanalysis.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdateTechniqueAnalysisRequest{
		Info: &techniqueanalysis,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertTechniqueAnalysis(t, resp.Info, &techniqueanalysis)
	}
}
