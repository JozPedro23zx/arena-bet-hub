package betentities

import (
	"errors"
	"strconv"
)

type simpleResult string
type overUnderResult string
type handCapAdvantage int

const (
	Win    simpleResult = "win"
	Draw   simpleResult = "draw"
	Defeat simpleResult = "defeat"
)

const (
	Over  overUnderResult = "over"
	Under overUnderResult = "under"
)

const (
	MoreOne  handCapAdvantage = 1
	MoreTwo  handCapAdvantage = 2
	MoreTree handCapAdvantage = 3
)

type TypeBet struct {
	simple      simpleResult
	overUnder   overUnderResult
	handCap     handCapAdvantage
	Proposition string
}

func (t *TypeBet) Simple() string {
	return string(t.simple)
}

func (t *TypeBet) OverUnder() string {
	return string(t.overUnder)
}

func (t *TypeBet) HandCap() int {
	return int(t.handCap)
}

func (t *TypeBet) SetSimple(predict string) error {
	predictSimple := simpleResult(predict)
	switch predictSimple {
	case Win, Draw, Defeat:
		t.simple = predictSimple
	default:
		return errors.New("invalid type bet")
	}
	return nil
}

func (t *TypeBet) SetOverUnder(predict string) error {
	predictOU := overUnderResult(predict)
	switch predictOU {
	case Over, Under:
		t.overUnder = overUnderResult(predict)
	default:
		return errors.New("invalid over/under result")
	}
	return nil
}

func (t *TypeBet) SetHandCap(predict string) error {
	predictNum, err := strconv.Atoi(predict)
	if err != nil {
		return errors.New("Bet Placed invalid. Hand Cap musb be integer")
	}
	predictHandCap := handCapAdvantage(predictNum)

	switch predictHandCap {
	case MoreOne, MoreTwo, MoreTree:
		t.handCap = handCapAdvantage(predictHandCap)
	default:
		return errors.New("invalid hand cap advantage")
	}

	return nil
}
