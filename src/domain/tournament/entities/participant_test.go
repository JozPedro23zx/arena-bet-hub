package tournament

import (
	"testing"

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
