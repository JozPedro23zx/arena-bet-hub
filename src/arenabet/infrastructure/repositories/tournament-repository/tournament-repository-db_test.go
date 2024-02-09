package tournamentrepository

import (
	"os"
	"testing"
	"time"

	Tournament "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-entities"
	"github.com/JozPedro23zx/arena-bet-hub/infrastructure/repositories/tournament-repository/fixture"
	"github.com/stretchr/testify/assert"
)

func TestTournamentDbInsert(t *testing.T) {
	migrationsDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)
	repository := NewTournamentRepositoryDB(db)

	location := Tournament.Location{
		Street:  "street",
		City:    "City",
		State:   "State",
		Country: "Country",
	}

	tournament := Tournament.NewTournament("tournament123", "Test tournament", time.Now(), location)

	err := repository.Insert(*tournament)

	assert.Nil(t, err)
}

func TestTournamentDbFind(t *testing.T) {
	migrationsDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)
	repository := NewTournamentRepositoryDB(db)

	location := Tournament.Location{
		Street:  "street",
		City:    "City",
		State:   "State",
		Country: "Country",
	}

	tournament := Tournament.NewTournament("tournament123", "Test tournament", time.Date(2024, time.March, 25, 22, 0, 0, 0, time.Local), location)

	err := repository.Insert(*tournament)
	assert.Nil(t, err)

	tournamentFound, err := repository.Find(tournament.ID)
	assert.Nil(t, err)
	assert.Equal(t, tournament, tournamentFound)
}

func TestTournamentDbUpdate(t *testing.T) {
	migrationsDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)
	repository := NewTournamentRepositoryDB(db)

	location := Tournament.Location{
		Street:  "street",
		City:    "City",
		State:   "State",
		Country: "Country",
	}
	tournament := Tournament.NewTournament("tournament123", "Test tournament", time.Date(2024, time.March, 25, 22, 0, 0, 0, time.Local), location)
	repository.Insert(*tournament)

	newLocation := Tournament.Location{
		Street:  "streetz",
		City:    "Cityz",
		State:   "Statez",
		Country: "Countryz",
	}
	newTournament := Tournament.NewTournament(tournament.ID, "Tournament updated", tournament.EventDate, newLocation)
	tournamentUpdated, err := repository.Update(*newTournament)

	assert.Nil(t, err)
	assert.Equal(t, tournamentUpdated, newTournament)
}
