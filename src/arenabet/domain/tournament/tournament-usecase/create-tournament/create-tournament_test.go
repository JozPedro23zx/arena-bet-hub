package createtournament

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	mock_tournament_repositories "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-repositories/mock"
	"github.com/golang/mock/gomock"

	Tournament "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-entities"
)

func TestCreateTournament(t *testing.T) {
	input := TournamentInputDto{
		ID:        "l12",
		Name:      "Namez",
		EventDate: time.Now(),
		Street:    "street",
		City:      "city",
		State:     "state",
		Country:   "country",
	}

	expectedOutput := TournamentOutputDto{
		ID:   "l12",
		Name: "Namez",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockTournamentRepository(ctrl)

	location := Tournament.Location{
		Street:  input.Street,
		City:    input.City,
		State:   input.State,
		Country: input.Country,
	}

	newTournament := Tournament.NewTournament(input.ID, input.Name, input.EventDate, location)

	repositoryMock.EXPECT().Find(input.ID)
	repositoryMock.EXPECT().Insert(*newTournament).Return(nil)

	usecase := NewCreateTournament(repositoryMock)

	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestTournamentAlreadyExist(t *testing.T) {
	input1 := TournamentInputDto{
		ID:        "l12",
		Name:      "Namez",
		EventDate: time.Now(),
		Street:    "street",
		City:      "city",
		State:     "state",
		Country:   "country",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockTournamentRepository(ctrl)

	location1 := Tournament.Location{
		Street:  input1.Street,
		City:    input1.City,
		State:   input1.State,
		Country: input1.Country,
	}

	tournament1 := Tournament.NewTournament(input1.ID, input1.Name, input1.EventDate, location1)

	repositoryMock.EXPECT().Find(input1.ID)
	repositoryMock.EXPECT().Insert(*tournament1).Return(nil)

	usecase := NewCreateTournament(repositoryMock)

	_, err := usecase.Execute(input1)
	assert.Nil(t, err)

	input2 := TournamentInputDto{
		ID:        "l12",
		Name:      "Battlez",
		EventDate: time.Now(),
		Street:    "street 2",
		City:      "city 2",
		State:     "state 3",
		Country:   "country 2",
	}

	expectedOutput := TournamentOutputDto{
		ID:   "l12",
		Name: "Namez",
	}

	repositoryMock.EXPECT().Find(input2.ID).Return(tournament1, nil)

	output, err := usecase.Execute(input2)

	assert.Error(t, err, "Tournament already exist")
	assert.Equal(t, expectedOutput, output)
}
