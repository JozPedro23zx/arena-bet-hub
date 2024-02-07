package findparticipant

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	Participant "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-entities"
	mock_tournament_repositories "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-repositories/mock"
)

func TestUpdateParticipant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockParticipantRepository(ctrl)

	input := ParcitipantInputDto{
		ID:            "1p12",
		Name:          "Luiz",
		NickName:      "Lipideo",
		CountryOrigin: "Spain",
	}

	participant := Participant.NewParticipant("1p12", "Luis", "Chronus", "USA")
	newParticipantUpdate := Participant.NewParticipant(input.ID, input.Name, input.NickName, input.CountryOrigin)

	repositoryMock.EXPECT().Find(input.ID).Return(participant, nil)
	repositoryMock.EXPECT().Update(*newParticipantUpdate).Return(newParticipantUpdate, nil)

	updateParticipant := NewUpdateParticipant(repositoryMock)
	output, err := updateParticipant.Execute(input)

	expectedOutput := ParticipantOutputDto{
		ID:            "1p12",
		Name:          "Luiz",
		NickName:      "Lipideo",
		CountryOrigin: "Spain",
		Tournmaents:   newParticipantUpdate.Tournamnets(),
	}

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestRegisterOnTournament(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockParticipantRepository(ctrl)

	input := TournamentInputDto{
		ParticipantID: "lp12",
		TournamentID:  "tournament123",
	}

	participant := Participant.NewParticipant("1p12", "Luis", "Chronus", "USA")
	newParticipantUpdate := Participant.NewParticipant(participant.ID, participant.Name, participant.NickName, participant.CountryOrigin)
	newParticipantUpdate.RegisterTournamentParticipation(input.TournamentID)

	repositoryMock.EXPECT().Find(input.ParticipantID).Return(participant, nil)
	repositoryMock.EXPECT().Update(*newParticipantUpdate).Return(newParticipantUpdate, nil)

	updateParticipant := NewUpdateParticipant(repositoryMock)
	output, err := updateParticipant.RegisterForTheTournament(input)

	expectedOutput := ParticipantOutputDto{
		ID:            "1p12",
		Name:          "Luis",
		NickName:      "Chronus",
		CountryOrigin: "USA",
		Tournmaents:   newParticipantUpdate.Tournamnets(),
	}

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestTournamentAlreadyExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockParticipantRepository(ctrl)

	input := TournamentInputDto{
		ParticipantID: "lp12",
		TournamentID:  "tournament123",
	}

	participant := Participant.NewParticipant("1p12", "Luis", "Chronus", "USA")
	participantAlreadyExist := participant.RegisterTournamentParticipation(input.TournamentID)

	repositoryMock.EXPECT().Find(input.ParticipantID).Return(participant, nil)

	updateParticipant := NewUpdateParticipant(repositoryMock)
	_, err := updateParticipant.RegisterForTheTournament(input)

	assert.Error(t, err, participantAlreadyExist)
}

func TestParticipantNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockParticipantRepository(ctrl)

	inputUpdate := ParcitipantInputDto{
		ID:            "1p12",
		Name:          "Luiz",
		NickName:      "Lipideo",
		CountryOrigin: "Spain",
	}

	participantNotFound := errors.New("Participant not found")

	repositoryMock.EXPECT().Find(inputUpdate.ID).Return(nil, participantNotFound)

	updateParticipant := NewUpdateParticipant(repositoryMock)
	_, err := updateParticipant.Execute(inputUpdate)

	assert.Error(t, err, "Participant not found")
}
