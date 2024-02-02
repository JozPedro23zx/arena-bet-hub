package updatetournament

import (
	"errors"
	"testing"
	"time"

	mock_tournament_repositories "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-repositories/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	CreateTournament "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-usecase/tournament/create-tournament"

	Tournament "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-entities"
)

func TestUpdateTournament(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockTournamentRepository(ctrl)

	input1 := CreateTournament.TournamentInputDto{
		ID:        "l12",
		Name:      "Namez",
		EventDate: time.Now(),
		Street:    "street",
		City:      "city",
		State:     "state",
		Country:   "country",
	}

	location := Tournament.Location{
		Street:  input1.Street,
		City:    input1.City,
		State:   input1.State,
		Country: input1.Country,
	}

	newTournament := Tournament.NewTournament(input1.ID, input1.Name, input1.EventDate, location)

	repositoryMock.EXPECT().Find(input1.ID)
	repositoryMock.EXPECT().Insert(*newTournament).Return(nil)

	createTournamentUC := CreateTournament.NewCreateTournament(repositoryMock)

	_, err := createTournamentUC.Execute(input1)

	input2 := TournamentInputDto{
		ID:        "l12",
		Name:      "Eventz",
		EventDate: input1.EventDate,
		Street:    "streetz",
		City:      "cityz",
		State:     "statez",
		Country:   "countryz",
	}

	expectedOutput := TournamentOutputDto{
		ID:        "l12",
		Name:      "Eventz",
		EventDate: input1.EventDate,
		Street:    "streetz",
		City:      "cityz",
		State:     "statez",
		Country:   "countryz",
		Finished:  false,
	}

	locationUpdated := Tournament.Location{
		Street:  input2.Street,
		City:    input2.City,
		State:   input2.State,
		Country: input2.Country,
	}

	newTournamentUpdated := Tournament.NewTournament(input2.ID, input2.Name, input2.EventDate, locationUpdated)

	repositoryMock.EXPECT().Find(input2.ID).Return(newTournament, nil)
	repositoryMock.EXPECT().Update(*newTournamentUpdated).Return(newTournamentUpdated, nil)

	updateTournamentUC := NewUpdateTournament(repositoryMock)

	output, err := updateTournamentUC.Execute(input2)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
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
