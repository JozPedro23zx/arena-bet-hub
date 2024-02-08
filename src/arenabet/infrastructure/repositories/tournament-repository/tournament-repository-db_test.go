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
