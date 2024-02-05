package updateresult

import (
	"time"

	tournament_entities "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-entities"
)

type OpenInputDto struct {
	ID   string `json:"id"`
	Open bool   `json:"open"`
}

type RankingInputDto struct {
	ID            string  `json:"id"`
	ParticipantID string  `json:"participantid"`
	Score         float64 `json:"score"`
}

type ResultOutputDto struct {
	ID           string                        `json:"id"`
	TurnamentID  string                        `json:"tournamentid"`
	Ranking      []tournament_entities.Ranking `json:"ranking"`
	Open         bool                          `json:"open"`
	DateFinished time.Time                     `json:"datefinished"`
}
