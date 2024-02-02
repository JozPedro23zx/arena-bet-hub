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

func TestUpdateParticipant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockParticipantRepository(ctrl)

	input := CreateParticipant.ParcitipantInputDto{
		ID:            "1p12",
		Name:          "Luis",
		NickName:      "Chronus",
		CountryOrigin: "USA",
	}

	participant := Participant.NewParticipant(input.ID, input.Name, input.NickName, input.CountryOrigin)

	repositoryMock.EXPECT().Find(input.ID).Return(nil, nil)
	repositoryMock.EXPECT().Insert(*participant).Return(nil)

	createParticipant := CreateParticipant.NewCreateParticipant(repositoryMock)
	_, err := createParticipant.Execute(input)

	inputUpdate := ParcitipantInputDto{
		ID:            "1p12",
		Name:          "Luiz",
		NickName:      "Lipideo",
		CountryOrigin: "Spain",
	}

	expectedOutput := ParticipantOutputDto{
		ID:            "1p12",
		Name:          "Luiz",
		NickName:      "Lipideo",
		CountryOrigin: "Spain",
	}

	newParticipantUpdate := Participant.NewParticipant(inputUpdate.ID, inputUpdate.Name, inputUpdate.NickName, inputUpdate.CountryOrigin)

	repositoryMock.EXPECT().Find(inputUpdate.ID).Return(participant, nil)
	repositoryMock.EXPECT().Update(*newParticipantUpdate).Return(newParticipantUpdate, nil)

	updateParticipant := NewUpdateParticipant(repositoryMock)
	output, err := updateParticipant.Execute(inputUpdate)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
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
