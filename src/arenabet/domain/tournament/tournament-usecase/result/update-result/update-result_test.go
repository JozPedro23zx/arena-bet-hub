package updateresult

import (
	"testing"

	mock_tournament_repositories "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-repositories/mock"
	mock_broker "github.com/JozPedro23zx/arena-bet-hub/infrastructure/broker/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	Tournament "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-entities"
)

func TestCloseResult(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockResultRepository(ctrl)
	producerMock := mock_broker.NewMockProducerInterface(ctrl)

	input := OpenInputDto{
		ID:   "result123",
		Open: false,
	}

	result := Tournament.NewResult(input.ID, "tournament123")
	resultUpdated := Tournament.NewResult(input.ID, "tournament123")
	resultUpdated.CloseResult()

	expectedOutput := ResultOutputDto{
		ID:           "result123",
		TurnamentID:  "tournament123",
		Ranking:      []RankingOutputDto{},
		Open:         false,
		DateFinished: resultUpdated.DateFinished,
	}

	repositoryMock.EXPECT().Find(input.ID).Return(result, nil)
	repositoryMock.EXPECT().Update(*resultUpdated).Return(resultUpdated, nil)
	producerMock.EXPECT().Publish(expectedOutput, []byte(input.ID), "close_open_result")

	updateResult := NewCloseOrOpenResult(repositoryMock, producerMock, "close_open_result")
	output, err := updateResult.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestRegisterParticipant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockResultRepository(ctrl)

	input := RankingInputDto{
		ResultID:      "result123",
		ParticipantID: "participant123",
		Score:         0,
	}

	result := Tournament.NewResult(input.ResultID, "tournament123")

	repositoryMock.EXPECT().Find(input.ResultID).Return(result, nil)

	updateResult := NewUpdateResult(repositoryMock)
	err := updateResult.AddParticipant(input)

	assert.Nil(t, err)
}

func TestParticipantAlreadyRegistered(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockResultRepository(ctrl)

	input := RankingInputDto{
		ResultID:      "result123",
		ParticipantID: "participant123",
		Score:         0,
	}

	result := Tournament.NewResult(input.ResultID, "tournament123")
	result.DefineRanking(input.ParticipantID)

	repositoryMock.EXPECT().Find(input.ResultID).Return(result, nil)

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
		ResultID:      "result123",
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

	rankingOutput := []RankingOutputDto{}
	for _, rank := range resultUpdated.Ranking() {
		output := RankingOutputDto{
			ParticipantID: rank.ParticipantId,
			Score:         rank.Score,
		}
		rankingOutput = append(rankingOutput, output)
	}

	expectedOutput := ResultOutputDto{
		ID:           "result123",
		TurnamentID:  "tournament123",
		Ranking:      rankingOutput,
		Open:         true,
		DateFinished: resultUpdated.DateFinished,
	}

	updateResult := NewUpdateResult(repositoryMock)
	output, err := updateResult.UpdateRanking(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
	assert.NotEmpty(t, output)
	assert.Equal(t, output.Ranking[0].ParticipantID, input.ParticipantID)
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
		ResultID:      "result123",
		ParticipantID: participant.ID,
		Score:         12.22,
	}

	result := Tournament.NewResult("result123", "tournament123")
	result.DefineRanking(participant.ID)
	result.UpdateRanking(participant.ID, 15.55)
	result.CloseResult()

	repositoryMock.EXPECT().Find(input.ResultID).Return(result, nil)

	updateResult := NewUpdateResult(repositoryMock)
	output, err := updateResult.UpdateRanking(input)

	rankingOutput := []RankingOutputDto{}
	for _, rank := range result.Ranking() {
		output := RankingOutputDto{
			ParticipantID: rank.ParticipantId,
			Score:         rank.Score,
		}
		rankingOutput = append(rankingOutput, output)
	}

	expectedOutput := ResultOutputDto{
		ID:           "result123",
		TurnamentID:  "tournament123",
		Ranking:      rankingOutput,
		Open:         false,
		DateFinished: result.DateFinished,
	}

	assert.Error(t, err, "Result has been closed")
	assert.Equal(t, expectedOutput, output)
}
