package tournament_entities

import (
	"errors"
	"sort"
	"time"
)

type Ranking struct {
	ParticipantId int
	Position      int
	Score         float64
}

type Result struct {
	ID           int
	TournamentId int
	Ranking      []Ranking
	Open         bool
	DateFinished time.Time
}

func NewResult(id int, tournamentId int) *Result {
	return &Result{
		ID:           id,
		TournamentId: tournamentId,
		Ranking:      []Ranking{},
		Open:         true,
		DateFinished: time.Time{},
	}
}

func (r *Result) CloseResult() {
	r.Open = false
	r.DateFinished = time.Now()
}

func (r *Result) DefineRanking(participant Participant) error {
	if r.Open {
		rank := Ranking{
			ParticipantId: participant.ID,
			Position:      0,
			Score:         0,
		}
		r.Ranking = append(r.Ranking, rank)
		return nil
	}
	return errors.New("Result cannot be changed")
}

func (r *Result) UpdateRanking(participantID int, newScore float64) error {
	if r.Open {
		for i, rank := range r.Ranking {
			if rank.ParticipantId == participantID {
				r.Ranking[i].Score = newScore

				sort.Slice(r.Ranking, func(i, j int) bool {
					return r.Ranking[i].Score > r.Ranking[j].Score
				})

				for i := range r.Ranking {
					r.Ranking[i].Position = i + 1
				}
				// return
			}
		}
		return nil
	}
	return errors.New("Result cannot be changed")
}
