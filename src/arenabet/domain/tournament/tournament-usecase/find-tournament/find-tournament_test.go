package findtournament

import (
	"errors"
	"testing"
	"time"

	mock_tournament_repositories "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-repositories/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	CreateTournament "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-usecase/create-tournament"

	Tournament "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-entities"
)

func TestFindTournament(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockTournamentRepository(ctrl)

	inputCreate := CreateTournament.TournamentInputDto{
		ID:        "l12",
		Name:      "Namez",
		EventDate: time.Now(),
		Street:    "street",
		City:      "city",
		State:     "state",
		Country:   "country",
	}

	location := Tournament.Location{
		Street:  inputCreate.Street,
		City:    inputCreate.City,
		State:   inputCreate.State,
		Country: inputCreate.Country,
	}

	newTournament := Tournament.NewTournament(inputCreate.ID, inputCreate.Name, inputCreate.EventDate, location)

	repositoryMock.EXPECT().Find(inputCreate.ID)
	repositoryMock.EXPECT().Insert(*newTournament).Return(nil)

	createTournament := CreateTournament.NewCreateTournament(repositoryMock)
	_, err := createTournament.Execute(inputCreate)

	inputFind := TournamentInputDto{ID: "l12"}
	expectedOutput := TournamentOutputDto{
		ID:        "l12",
		Name:      "Namez",
		EventDate: inputCreate.EventDate,
		Street:    "street",
		City:      "city",
		State:     "state",
		Country:   "country",
		Finished:  false,
	}

	repositoryMock.EXPECT().Find(inputFind.ID).Return(newTournament, nil)
	findTournament := NewFindTournament(repositoryMock)
	output, err := findTournament.Execute(inputFind)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestTournamentNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockTournamentRepository(ctrl)

	inputFind := TournamentInputDto{ID: "l12"}

	tournamentNotFoundErr := errors.New("Tournament not found")

	repositoryMock.EXPECT().Find(inputFind.ID).Return(nil, tournamentNotFoundErr)
	findTournament := NewFindTournament(repositoryMock)
	_, err := findTournament.Execute(inputFind)

	assert.Error(t, err, "Tournament not found")
}
