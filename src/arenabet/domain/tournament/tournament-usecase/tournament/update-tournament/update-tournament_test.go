package updatetournament

import (
	"errors"
	"testing"
	"time"

	mock_tournament_repositories "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-repositories/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	Tournament "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-entities"
)

func TestUpdateTournament(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockTournamentRepository(ctrl)

	location := Tournament.Location{
		Street:  "street",
		City:    "city",
		State:   "state",
		Country: "country",
	}

	newTournament := Tournament.NewTournament("l12", "Namez", time.Now(), location)

	input := TournamentInputDto{
		ID:        "l12",
		Name:      "Eventz",
		EventDate: newTournament.EventDate,
		Street:    "streetz",
		City:      "cityz",
		State:     "statez",
		Country:   "countryz",
	}

	expectedOutput := TournamentOutputDto{
		ID:           "l12",
		Name:         "Eventz",
		EventDate:    newTournament.EventDate,
		Street:       "streetz",
		City:         "cityz",
		State:        "statez",
		Country:      "countryz",
		Participants: newTournament.Participants(),
		Finished:     false,
	}

	locationUpdated := Tournament.Location{
		Street:  input.Street,
		City:    input.City,
		State:   input.State,
		Country: input.Country,
	}

	newTournamentUpdated := Tournament.NewTournament(input.ID, input.Name, input.EventDate, locationUpdated)

	repositoryMock.EXPECT().Find(input.ID).Return(newTournament, nil)
	repositoryMock.EXPECT().Update(*newTournamentUpdated).Return(newTournamentUpdated, nil)

	updateTournamentUC := NewUpdateTournament(repositoryMock)
	output, err := updateTournamentUC.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestAddParticipant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockTournamentRepository(ctrl)

	input := ParticipantListDto{
		IDTournament:  "tournament123",
		ParticipantID: "participant123",
		Add:           true,
	}

	location := Tournament.Location{
		Street:  "street",
		City:    "city",
		State:   "state",
		Country: "country",
	}

	newTournament := Tournament.NewTournament(input.IDTournament, "event12", time.Now(), location)
	updatedTornament := Tournament.NewTournament(newTournament.ID, newTournament.Name, newTournament.EventDate, newTournament.Location)
	updatedTornament.RegisterParticipant(input.ParticipantID)

	repositoryMock.EXPECT().Find(input.IDTournament).Return(newTournament, nil)
	repositoryMock.EXPECT().Update(*updatedTornament).Return(updatedTornament, nil)

	updateTournament := NewUpdateTournament(repositoryMock)
	output, err := updateTournament.AddParticipant(input)

	assert.Nil(t, err)
	assert.Equal(t, output.Participants[0], input.ParticipantID)
}

func TestRemoveParticipant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockTournamentRepository(ctrl)

	input := ParticipantListDto{
		IDTournament:  "tournament123",
		ParticipantID: "participant2",
		Add:           false,
	}

	location := Tournament.Location{
		Street:  "street",
		City:    "city",
		State:   "state",
		Country: "country",
	}

	newTournament := Tournament.NewTournament(input.IDTournament, "event12", time.Now(), location)
	newTournament.RegisterParticipant("participant1")
	newTournament.RegisterParticipant("participant2")
	newTournament.RegisterParticipant("participant3")

	updatedTornament := Tournament.NewTournament(newTournament.ID, newTournament.Name, newTournament.EventDate, newTournament.Location)
	updatedTornament.RegisterParticipant("participant1")
	updatedTornament.RegisterParticipant("participant3")

	repositoryMock.EXPECT().Find(input.IDTournament).Return(newTournament, nil)
	repositoryMock.EXPECT().Update(*updatedTornament).Return(updatedTornament, nil)

	updateTournament := NewUpdateTournament(repositoryMock)
	output, err := updateTournament.AddParticipant(input)

	assert.Nil(t, err)
	assert.Equal(t, output.Participants[0], "participant1")
	assert.Equal(t, output.Participants[1], "participant3")
}

func TestParticipantAlreadyExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockTournamentRepository(ctrl)

	input := ParticipantListDto{
		IDTournament:  "tournament123",
		ParticipantID: "participant123",
		Add:           true,
	}

	location := Tournament.Location{
		Street:  "street",
		City:    "city",
		State:   "state",
		Country: "country",
	}

	newTournament := Tournament.NewTournament(input.IDTournament, "event12", time.Now(), location)
	participantAlreadyExist := newTournament.RegisterParticipant(input.ParticipantID)

	repositoryMock.EXPECT().Find(input.IDTournament).Return(newTournament, nil)

	updateTournament := NewUpdateTournament(repositoryMock)
	_, err := updateTournament.AddParticipant(input)

	assert.Error(t, err, participantAlreadyExist)
}

func TestTournamentNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockTournamentRepository(ctrl)

	input := TournamentInputDto{
		ID:        "l12",
		Name:      "Eventz",
		EventDate: time.Now(),
		Street:    "streetz",
		City:      "cityz",
		State:     "statez",
		Country:   "countryz",
	}

	tournamentNotFoundErr := errors.New("Tournament not found")

	repositoryMock.EXPECT().Find(input.ID).Return(nil, tournamentNotFoundErr)

	updateTournamentUC := NewUpdateTournament(repositoryMock)

	_, err := updateTournamentUC.Execute(input)

	assert.Error(t, err, "Tournament not found")
}
