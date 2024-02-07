package findtournament

import (
	"errors"
	"testing"
	"time"

	mock_tournament_repositories "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-repositories/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	Tournament "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-entities"
)

func TestFindTournament(t *testing.T) {
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

	inputFind := TournamentInputDto{ID: "l12"}
	expectedOutput := TournamentOutputDto{
		ID:           "l12",
		Name:         "Namez",
		EventDate:    newTournament.EventDate,
		Street:       "street",
		City:         "city",
		State:        "state",
		Country:      "country",
		Participants: newTournament.Participants(),
		Finished:     false,
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
