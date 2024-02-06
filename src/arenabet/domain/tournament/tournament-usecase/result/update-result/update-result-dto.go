package updateresult

import "time"

type OpenInputDto struct {
	ID   string `json:"id"`
	Open bool   `json:"open"`
}

type RankingInputDto struct {
	ResultID      string  `json:"id"`
	ParticipantID string  `json:"participantid"`
	Score         float64 `json:"score"`
}

type RankingOutputDto struct {
	ParticipantID string  `json:"participantid"`
	Score         float64 `json:"score"`
}

type ResultOutputDto struct {
	ID           string             `json:"id"`
	TurnamentID  string             `json:"tournamentid"`
	Ranking      []RankingOutputDto `json:"ranking"`
	Open         bool               `json:"open"`
	DateFinished time.Time          `json:"datefinished"`
}
