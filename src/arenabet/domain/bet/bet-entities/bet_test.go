package betentities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateAward(t *testing.T) {
	bet := NewBet("1", "user1", "tournament1", 100.0)
	odds := 2.0
	expectedPayment := 200.0

	payment := bet.CalculateAward(odds, 0)

	assert.Equal(t, expectedPayment, payment)
}

func TestDefineType(t *testing.T) {
	bet := NewBet("1", "user1", "tournament1", 100.0)
	err := bet.DefineType("simple", "win")

	assert.Nil(t, err)
	assert.Equal(t, bet.Type.simple, Win)
}

func TestVerifyBetResult(t *testing.T) {
	bet := NewBet("1", "user1", "tournament1", 100.0)
	bet.DefineType("simple", "win")

	result := "win"
	odds := 2.0
	oddsResult := bet.VerifyBetResult(result, odds)

	assert.Equal(t, odds, oddsResult)

	bet2 := NewBet("2", "user2", "tournament2", 100.0)
	bet2.DefineType("simple", "win")

	result = "defeat"
	odds = 2.0
	oddsResult = bet.VerifyBetResult(result, odds)

	assert.NotEqual(t, odds, oddsResult)
	assert.Equal(t, 0.0, oddsResult)
}
