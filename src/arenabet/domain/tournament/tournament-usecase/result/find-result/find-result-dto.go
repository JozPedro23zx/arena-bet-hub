package findresult

type ResultInputDto struct {
	ID string `json:"id"`
}

type ResultOutputDto struct {
	ID          string `json:"id"`
	TurnamentID string `json:"tournamentid"`
	Open        bool   `json:"open"`
}
