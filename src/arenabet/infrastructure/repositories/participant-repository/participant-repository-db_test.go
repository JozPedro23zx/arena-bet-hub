package participantrepository

import (
	"os"
	"testing"

	Tournament "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-entities"
	"github.com/JozPedro23zx/arena-bet-hub/infrastructure/repositories/fixture"
	"github.com/stretchr/testify/assert"
)

func TestParticipantInsertDb(t *testing.T) {
	migrationsDir := os.DirFS("../fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)
	repository := NewParticipantRepositoryDB(db)

	participant := Tournament.NewParticipant("participant123", "Joe", "Jo3fighter", "USA")
	err := repository.Insert(*participant)

	assert.Nil(t, err)
}

func TestParticipantFindDb(t *testing.T) {
	migrationsDir := os.DirFS("../fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)
	repository := NewParticipantRepositoryDB(db)

	participant := Tournament.NewParticipant("participant123", "Joe", "Jo3fighter", "USA")
	_ = repository.Insert(*participant)

	participantFound, err := repository.Find(participant.ID)

	assert.Nil(t, err)
	assert.Equal(t, participant, participantFound)
}

func TestParticipantUpdateDb(t *testing.T) {
	migrationsDir := os.DirFS("../fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)
	repository := NewParticipantRepositoryDB(db)

	participant := Tournament.NewParticipant("participant123", "Joe", "Jo3fighter", "USA")
	_ = repository.Insert(*participant)

	newParticipant := Tournament.NewParticipant(participant.ID, "Crhis Joe", "Jo3Z", participant.CountryOrigin)
	participantUpdated, err := repository.Update(*newParticipant)

	assert.Nil(t, err)
	assert.Equal(t, newParticipant, participantUpdated)
	assert.NotEqual(t, participant, participantUpdated)
}
