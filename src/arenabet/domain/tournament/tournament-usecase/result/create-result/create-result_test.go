package createresult

import (
	"errors"
	"testing"

	mock_tournament_repositories "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-repositories/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	Tournament "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-entities"
)

func TestCreateResult(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockResultRepository(ctrl)

	input := ResultInputDto{
		ID:          "result123",
		TurnamentID: "tournament123",
	}

	expectedOutput := ResultOutputDto{
		ID:          "result123",
		TurnamentID: "tournament123",
		Open:        true,
	}

	newResult := Tournament.NewResult(input.ID, input.TurnamentID)

	repositoryMock.EXPECT().Find(input.ID).Return(nil, nil)
	repositoryMock.EXPECT().Insert(*newResult).Return(nil)

	createResult := NewCreateResult(repositoryMock)
	output, err := createResult.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestResultAlreadyExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockResultRepository(ctrl)

	input := ResultInputDto{
		ID:          "result123",
		TurnamentID: "tournament123",
	}

	expectedOutput := ResultOutputDto{
		ID:          "result123",
		TurnamentID: "tournament123",
		Open:        true,
	}

	newResult := Tournament.NewResult(input.ID, input.TurnamentID)
	repositoryMock.EXPECT().Find(input.ID).Return(newResult, nil)

	createResult := NewCreateResult(repositoryMock)
	output, err := createResult.Execute(input)

	resultAlreadyExist := errors.New("Result alredy existe")

	assert.Error(t, err, resultAlreadyExist)
	assert.Equal(t, expectedOutput, output)
}
