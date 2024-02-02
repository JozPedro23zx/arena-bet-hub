package createparticipant

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	Participant "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-entities"
	mock_tournament_repositories "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-repositories/mock"
)

func TestCreateParticipant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockParticipantRepository(ctrl)

	input := ParcitipantInputDto{
		ID:            "1p12",
		Name:          "Luis",
		NickName:      "Chronus",
		CountryOrigin: "USA",
	}

	expectedOutput := ParticipantOutputDto{
		ID:       "1p12",
		NickName: "Chronus",
	}

	participant := Participant.NewParticipant(input.ID, input.Name, input.NickName, input.CountryOrigin)

	repositoryMock.EXPECT().Find(input.ID).Return(nil, nil)
	repositoryMock.EXPECT().Insert(*participant).Return(nil)

	createParticipant := NewCreateParticipant(repositoryMock)
	output, err := createParticipant.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestParticipantAlreadyExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockParticipantRepository(ctrl)

	input1 := ParcitipantInputDto{
		ID:            "1p12",
		Name:          "Luis",
		NickName:      "Chronus",
		CountryOrigin: "USA",
	}

	participant := Participant.NewParticipant(input1.ID, input1.Name, input1.NickName, input1.CountryOrigin)

	repositoryMock.EXPECT().Find(input1.ID).Return(nil, nil)
	repositoryMock.EXPECT().Insert(*participant).Return(nil)

	createParticipant := NewCreateParticipant(repositoryMock)
	_, err := createParticipant.Execute(input1)

	input2 := ParcitipantInputDto{
		ID:            "1p12",
		Name:          "Luis",
		NickName:      "Fenrrir",
		CountryOrigin: "USA",
	}

	expectedOutput := ParticipantOutputDto{
		ID:       "1p12",
		NickName: "Chronus",
	}

	repositoryMock.EXPECT().Find(input2.ID).Return(participant, nil)

	output, err := createParticipant.Execute(input2)

	assert.Error(t, err, "Participant already registered")
	assert.Equal(t, expectedOutput, output)
}
