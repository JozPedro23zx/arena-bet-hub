package betentities

import (
	"fmt"
	"strconv"
)

type Bet struct {
	ID           string
	IDUser       string
	IDTournament string
	Type         TypeBet
	Value        float64
}

func NewBet(id string, user string, tournament string, value float64) *Bet {
	return &Bet{ID: id, IDUser: user, IDTournament: tournament, Value: value}
}

func (b *Bet) DefineType(typeDefined string, value string) error {
	switch typeDefined {
	case "simple":
		err := b.Type.SetSimple(value)
		if err != nil {
			return err
		}
	case "handcap":
		err := b.Type.SetHandCap(value)
		if err != nil {
			return err
		}
	case "overunder":
		err := b.Type.SetOverUnder(value)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Bet) CalculateAward(odds float64, bonus float64) float64 {
	payment := b.Value*odds + bonus
	return payment
}

func (tb *Bet) VerifyBetResult(result string, odds float64) float64 {
	switch {
	case tb.Type.simple != "":
		if tb.Type.simple == simpleResult(result) {
			return odds
		} else {
			return 0
		}
	case tb.Type.overUnder != "":
		if tb.Type.overUnder == overUnderResult(result) {
			return odds
		} else {
			return 0
		}
	case tb.Type.handCap > 0 && tb.Type.handCap < 3:
		resultNum, err := strconv.Atoi(result)
		if err != nil {
			fmt.Println("Bet Placed invalid. Hand Cap error", err)
			return 0
		}
		if tb.Type.handCap == handCapAdvantage(resultNum) {
			return odds
		} else {
			return 0
		}
	}

	return 0
}
