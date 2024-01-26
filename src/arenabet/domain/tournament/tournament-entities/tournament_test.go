package tournament_entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTournament(t *testing.T) {
	location := Location{
		Street:  "street",
		City:    "City",
		State:   "State",
		Country: "Country",
	}

	tournament := NewTournament(1, "Tournament test", time.Date(2024, time.March, 25, 22, 0, 0, 0, time.UTC), location)

	assert.Equal(t, 1, tournament.ID)
	assert.Equal(t, "Tournament test", tournament.Name)
	assert.NotZero(t, tournament.EventDate)
	assert.Equal(t, time.Date(2024, time.March, 25, 22, 0, 0, 0, time.UTC), tournament.EventDate)
	assert.Equal(t, location, tournament.Location)
	assert.Equal(t, []int{}, tournament.IDParticipants)
}

func TestUpdateTournament(t *testing.T) {
	location := Location{
		Street:  "street",
		City:    "City",
		State:   "State",
		Country: "Country",
	}

	tournament := NewTournament(1, "Tournament test", time.Date(2024, time.March, 25, 22, 0, 0, 0, time.UTC), location)

	updatedName := "Tournament test updated"
	updatedTime := time.Date(2024, time.April, 25, 22, 0, 0, 0, time.UTC)
	updatedLocation := Location{
		Street:  "Street Updated",
		City:    "City Updated",
		State:   "State Updated",
		Country: "Country Updated",
	}

	tournament.UpdateTournament(tournament.Name, tournament.EventDate, updatedLocation)

	assert.Equal(t, "Tournament test", tournament.Name)
	assert.NotZero(t, tournament.EventDate)
	assert.Equal(t, updatedLocation, tournament.Location)

	tournament.UpdateTournament(tournament.Name, updatedTime, updatedLocation)

	assert.Equal(t, "Tournament test", tournament.Name)
	assert.NotZero(t, tournament.EventDate)
	assert.Equal(t, updatedTime, tournament.EventDate)
	assert.Equal(t, updatedLocation, tournament.Location)

	tournament.UpdateTournament(updatedName, updatedTime, updatedLocation)

	assert.Equal(t, updatedName, tournament.Name)
	assert.NotZero(t, tournament.EventDate)
	assert.Equal(t, updatedLocation, tournament.Location)

	tournament.FinishTournament()

	updatedName2 := "Tournament test updated for the second time"
	err := tournament.UpdateTournament(updatedName, tournament.EventDate, updatedLocation)

	assert.NotEqual(t, updatedName2, tournament.Name)
	assert.Equal(t, updatedName, tournament.Name)
	assert.Error(t, err, "This tournament has finished")
}

func TestRegisterParticipant(t *testing.T) {
	location := Location{
		Street:  "street",
		City:    "City",
		State:   "State",
		Country: "Country",
	}
	tournament := NewTournament(1, "Tournament test", time.Date(2024, time.March, 25, 22, 0, 0, 0, time.UTC), location)
	participant := NewParticipant(1, "Jon", "Jon90z", "USA")

	tournament.RegisterParticipant(*participant)

	assert.NotEmpty(t, tournament.IDParticipants)
	assert.Equal(t, participant.ID, tournament.IDParticipants[0])
}

func TestRemoveParticipant(t *testing.T) {
	location := Location{
		Street:  "street",
		City:    "City",
		State:   "State",
		Country: "Country",
	}
	tournament := NewTournament(1, "Tournament test", time.Date(2024, time.March, 25, 22, 0, 0, 0, time.UTC), location)
	participant := NewParticipant(1, "Jon", "Jon90z", "USA")
	participant2 := NewParticipant(2, "Mateus", "Mateuzin10", "Brazil")

	tournament.RegisterParticipant(*participant)

	assert.NotEmpty(t, tournament.IDParticipants)
	assert.Equal(t, participant.ID, tournament.IDParticipants[0])

	tournament.RemoveParticipant(participant.ID)
	assert.Empty(t, tournament.IDParticipants)

	tournament.FinishTournament()

	err := tournament.RegisterParticipant(*participant2)

	assert.Empty(t, tournament.IDParticipants)
	assert.Error(t, err, "This tournament has finished")
}
