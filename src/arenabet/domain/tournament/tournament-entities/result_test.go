package tournament_entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewResult(t *testing.T) {

	location := Location{
		Street:  "street",
		City:    "City",
		State:   "State",
		Country: "Country",
	}

	tournament := NewTournament(1, "Tournament test", time.Date(2024, time.March, 25, 22, 0, 0, 0, time.UTC), location)

	result := NewResult(1, tournament.ID)

	assert.Equal(t, 1, result.ID)
	assert.Equal(t, tournament.ID, result.ID)
}

func TestDefineRanking(t *testing.T) {
	location := Location{
		Street:  "street",
		City:    "City",
		State:   "State",
		Country: "Country",
	}
	tournament := NewTournament(1, "Tournament test", time.Date(2024, time.March, 25, 22, 0, 0, 0, time.UTC), location)

	participant := NewParticipant(1, "Arthur", "Excalibur", "USA")
	participant2 := NewParticipant(2, "Jon", "Jon90z", "USA")

	result := NewResult(1, tournament.ID)
	err := result.DefineRanking((*participant))

	assert.NoError(t, err)
	assert.NotEmpty(t, result.Ranking)

	result.CloseResult()
	err = result.DefineRanking(*participant2)

	assert.Error(t, err, "Result cannot be changed")
}

func TestUpdateRanking(t *testing.T) {
	location := Location{
		Street:  "street",
		City:    "City",
		State:   "State",
		Country: "Country",
	}

	tournament := NewTournament(1, "Tournament test", time.Date(2024, time.March, 25, 22, 0, 0, 0, time.UTC), location)

	participant1 := NewParticipant(1, "Arthur", "Excalibur", "USA")
	participant2 := NewParticipant(2, "Jon", "Jon90z", "USA")
	participant3 := NewParticipant(3, "Mateus", "Mateuzin10", "Brazil")

	result := NewResult(1, tournament.ID)

	result.DefineRanking(*participant1)
	result.DefineRanking(*participant2)
	result.DefineRanking(*participant3)

	score1 := 10.11
	socre2 := 24.14
	score3 := 0.28
	score4 := 1.45

	result.UpdateRanking(participant1.ID, score1)
	result.UpdateRanking(participant2.ID, socre2)
	result.UpdateRanking(participant3.ID, score3)

	assert.Equal(t, 1, result.Ranking[0].Position)
	assert.Equal(t, participant2.ID, result.Ranking[0].ParticipantId)

	assert.Equal(t, 2, result.Ranking[1].Position)
	assert.Equal(t, participant1.ID, result.Ranking[1].ParticipantId)

	assert.Equal(t, 3, result.Ranking[2].Position)
	assert.Equal(t, participant3.ID, result.Ranking[2].ParticipantId)

	result.CloseResult()

	err := result.UpdateRanking(participant1.ID, score4)

	assert.Error(t, err, "Result cannot be changed")
}
