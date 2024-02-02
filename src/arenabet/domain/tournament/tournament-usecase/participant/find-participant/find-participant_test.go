package findparticipant

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	Participant "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-entities"
	mock_tournament_repositories "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-repositories/mock"
	CreateParticipant "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-usecase/participant/create-participant"
)

func TestFindParticipant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockParticipantRepository(ctrl)

	input1 := CreateParticipant.ParcitipantInputDto{
		ID:            "1p12",
		Name:          "Luis",
		NickName:      "Chronus",
		CountryOrigin: "USA",
	}

	expectedOutput := ParticipantOutputDto{
		ID:            "1p12",
		Name:          "Luis",
		NickName:      "Chronus",
		CountryOrigin: "USA",
	}

	participant := Participant.NewParticipant(input1.ID, input1.Name, input1.NickName, input1.CountryOrigin)

	repositoryMock.EXPECT().Find(input1.ID).Return(nil, nil)
	repositoryMock.EXPECT().Insert(*participant).Return(nil)

	createParticipant := CreateParticipant.NewCreateParticipant(repositoryMock)
	_, err := createParticipant.Execute(input1)

	input2 := ParcitipantInputDto{
		ID: "1p12",
	}

	repositoryMock.EXPECT().Find(input2.ID).Return(participant, nil)

	findParticipant := NewFindParticipant(repositoryMock)

	output, err := findParticipant.Execute(input2)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestParticipantNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockParticipantRepository(ctrl)

	input := ParcitipantInputDto{
		ID: "L12",
	}

	participantNotFound := errors.New("Participant not found")

	repositoryMock.EXPECT().Find(input.ID).Return(nil, participantNotFound)

	findParticipant := NewFindParticipant(repositoryMock)

	_, err := findParticipant.Execute(input)

	assert.Error(t, err, "Participant not found")
}
