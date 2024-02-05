package findresult

import (
	"errors"
	"testing"

	mock_tournament_repositories "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-repositories/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	Tournament "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-entities"
)

func TestFindResult(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockResultRepository(ctrl)

	input := ResultInputDto{
		ID: "result123",
	}
	expectedOutput := ResultOutputDto{
		ID:          "result123",
		TurnamentID: "tournament123",
		Open:        true,
	}

	result := Tournament.NewResult("result123", "tournament123")

	repositoryMock.EXPECT().Find(input.ID).Return(result, nil)

	findResult := NewFindResult(repositoryMock)

	output, err := findResult.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestResultNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockResultRepository(ctrl)

	input := ResultInputDto{
		ID: "result123",
	}

	resultNotFound := errors.New("Result not found")

	repositoryMock.EXPECT().Find(input.ID).Return(nil, resultNotFound)

	findResult := NewFindResult(repositoryMock)

	_, err := findResult.Execute(input)

	assert.Error(t, err, "Result not found")
}
