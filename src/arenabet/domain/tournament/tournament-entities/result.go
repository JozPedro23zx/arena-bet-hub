package tournament_entities

import (
	"errors"
	"sort"
	"time"
)

type Ranking struct {
	ParticipantId string
	Position      int
	Score         float64
}

type Result struct {
	ID           string
	TournamentId string
	ranking      []Ranking
	Open         bool
	DateFinished time.Time
}

func NewResult(id string, tournamentId string) *Result {
	return &Result{
		ID:           id,
		TournamentId: tournamentId,
		ranking:      []Ranking{},
		Open:         true,
		DateFinished: time.Time{},
	}
}

func (r *Result) Ranking() []Ranking {
	return r.ranking
}

func (r *Result) CloseResult() {
	r.Open = false
	r.DateFinished = time.Now().Truncate(time.Millisecond)
}

func (r *Result) DefineRanking(participantId string) (Ranking, error) {
	if r.Open {
		for _, existingRanking := range r.ranking {
			if existingRanking.ParticipantId == participantId {

				return existingRanking, errors.New("participant already registered")
			}
		}

		rank := Ranking{
			ParticipantId: participantId,
			Position:      0,
			Score:         0,
		}
		r.ranking = append(r.ranking, rank)
		return rank, nil
	}
	return Ranking{}, errors.New("Result cannot be changed")
}

func (r *Result) UpdateRanking(participantID string, newScore float64) error {
	if r.Open {
		for i, rank := range r.ranking {
			if rank.ParticipantId == participantID {
				r.ranking[i].Score = newScore

				sort.Slice(r.ranking, func(i, j int) bool {
					return r.ranking[i].Score > r.ranking[j].Score
				})

				for i := range r.ranking {
					r.ranking[i].Position = i + 1
				}
				// return
			}
		}
		return nil
	}
	return errors.New("Result cannot be changed")
}
