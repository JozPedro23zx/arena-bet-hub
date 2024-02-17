package createtournament

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	mock_tournament_repositories "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-repositories/mock"
	mock_broker "github.com/JozPedro23zx/arena-bet-hub/infrastructure/broker/mock"
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
		ID:       "l12",
		Name:     "Namez",
		Finished: false,
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

	producerMock := mock_broker.NewMockProducerInterface(ctrl)
	producerMock.EXPECT().Publish(expectedOutput, []byte(input.ID), "tournament_created")

	usecase := NewCreateTournament(repositoryMock, producerMock, "tournament_created")

	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestTournamentAlreadyExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_tournament_repositories.NewMockTournamentRepository(ctrl)
	producerMock := mock_broker.NewMockProducerInterface(ctrl)

	location1 := Tournament.Location{
		Street:  "street",
		City:    "city",
		State:   "state",
		Country: "country",
	}
	tournament := Tournament.NewTournament("l12", "Namez", time.Now(), location1)

	input := TournamentInputDto{
		ID:        "l12",
		Name:      "Battlez",
		EventDate: tournament.EventDate,
		Street:    "street 2",
		City:      "city 2",
		State:     "state 3",
		Country:   "country 2",
	}

	expectedOutput := TournamentOutputDto{
		ID:       "l12",
		Name:     "Namez",
		Finished: false,
	}

	repositoryMock.EXPECT().Find(input.ID).Return(tournament, nil)
	usecase := NewCreateTournament(repositoryMock, producerMock, "tournament_created")

	output, err := usecase.Execute(input)

	assert.Error(t, err, "tournament already exist")
	assert.Equal(t, expectedOutput, output)
}
