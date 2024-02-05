package updateresult

import (
	"testing"

	mock_tournament_repositories "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-repositories/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	Tournament "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-entities"
)

func TestCloseResult(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockResultRepository(ctrl)

	input := OpenInputDto{
		ID:   "result123",
		Open: false,
	}

	result := Tournament.NewResult(input.ID, "tournament123")
	resultUpdated := Tournament.NewResult(input.ID, "tournament123")
	resultUpdated.CloseResult()

	repositoryMock.EXPECT().Find(input.ID).Return(result, nil)
	repositoryMock.EXPECT().Update(*resultUpdated).Return(resultUpdated, nil)

	updateResult := NewUpdateResult(repositoryMock)
	output, err := updateResult.CloseOrOpenResult(input)

	expectedOutput := ResultOutputDto{
		ID:           "result123",
		TurnamentID:  "tournament123",
		Ranking:      []Tournament.Ranking{},
		Open:         false,
		DateFinished: resultUpdated.DateFinished,
	}

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestRegisterParticipant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockResultRepository(ctrl)

	input := RankingInputDto{
		ID:            "result123",
		ParticipantID: "participant123",
		Score:         0,
	}

	result := Tournament.NewResult(input.ID, "tournament123")

	repositoryMock.EXPECT().Find(input.ID).Return(result, nil)

	updateResult := NewUpdateResult(repositoryMock)
	err := updateResult.AddParticipant(input)

	assert.Nil(t, err)
}

func TestParticipantAlreadyRegistered(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockResultRepository(ctrl)

	input := RankingInputDto{
		ID:            "result123",
		ParticipantID: "participant123",
		Score:         0,
	}

	result := Tournament.NewResult(input.ID, "tournament123")
	result.DefineRanking(input.ParticipantID)

	repositoryMock.EXPECT().Find(input.ID).Return(result, nil)

	updateResult := NewUpdateResult(repositoryMock)
	err := updateResult.AddParticipant(input)

	assert.Error(t, err, "participant already registered")
}

func TestUpdateRanking(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockResultRepository(ctrl)

	participant := Tournament.Participant{
		ID:            "participant123",
		Name:          "Wesker",
		NickName:      "Krauzer",
		CountryOrigin: "USA",
	}

	input := RankingInputDto{
		ID:            "result123",
		ParticipantID: participant.ID,
		Score:         12.22,
	}

	result := Tournament.NewResult("result123", "tournament123")
	result.DefineRanking(participant.ID)

	resultUpdated := Tournament.NewResult("result123", "tournament123")
	resultUpdated.DefineRanking(participant.ID)
	resultUpdated.UpdateRanking(input.ParticipantID, input.Score)

	repositoryMock.EXPECT().Find(result.ID).Return(result, nil)
	repositoryMock.EXPECT().Update(*resultUpdated).Return(resultUpdated, nil)

	expectedOutput := ResultOutputDto{
		ID:           "result123",
		TurnamentID:  "tournament123",
		Ranking:      resultUpdated.Ranking,
		Open:         true,
		DateFinished: resultUpdated.DateFinished,
	}

	updateResult := NewUpdateResult(repositoryMock)
	output, err := updateResult.UpdateRanking(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
	assert.NotEmpty(t, output)
	assert.Equal(t, output.Ranking[0].ParticipantId, input.ParticipantID)
}

func TestResultHasBeenClosed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockResultRepository(ctrl)

	participant := Tournament.Participant{
		ID:            "participant123",
		Name:          "Wesker",
		NickName:      "Krauzer",
		CountryOrigin: "USA",
	}

	input := RankingInputDto{
		ID:            "result123",
		ParticipantID: participant.ID,
		Score:         12.22,
	}

	result := Tournament.NewResult("result123", "tournament123")
	result.DefineRanking(participant.ID)
	result.UpdateRanking(participant.ID, 15.55)
	result.CloseResult()

	repositoryMock.EXPECT().Find(input.ID).Return(result, nil)

	updateResult := NewUpdateResult(repositoryMock)
	output, err := updateResult.UpdateRanking(input)

	expectedOutput := ResultOutputDto{
		ID:           "result123",
		TurnamentID:  "tournament123",
		Ranking:      result.Ranking,
		Open:         false,
		DateFinished: result.DateFinished,
	}

	assert.Error(t, err, "Result has been closed")
	assert.Equal(t, expectedOutput, output)
}
