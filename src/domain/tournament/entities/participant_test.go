package tournament

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewParticipant(t *testing.T) {
	participant := NewParticipant(1, "Arthur", "Excalibur", "USA")

	assert.Equal(t, "Arthur", participant.Name)
	assert.Equal(t, "Excalibur", participant.NickName)
	assert.Equal(t, "USA", participant.CountryOrigin)
}

func TestUpdateParticipant(t *testing.T) {
	participant := NewParticipant(1, "Arthur", "Excalibur", "USA")

	updatedNick := "Excalibur_t286"

	participant.UpdateParticipant(participant.Name, updatedNick, participant.CountryOrigin)
}

func TestRegisterTournamentParticipation(t *testing.T) {
	location := Location{
		Street:  "street",
		City:    "City",
		State:   "State",
		Country: "Country",
	}
	tournament := NewTournament(1, "Tournament test", time.Date(2024, time.March, 25, 22, 0, 0, 0, time.UTC), location)

	participant := NewParticipant(2, "Jon", "Jon90z", "USA")
	participant.RegisterTournamentParticipation(*tournament)

	assert.NotEmpty(t, participant.Tournaments)
	assert.Equal(t, participant.Tournaments[0], tournament.ID)
}
