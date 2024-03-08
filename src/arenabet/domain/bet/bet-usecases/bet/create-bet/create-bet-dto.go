package createbet

type BetInputDto struct {
	ID           string  `json:"id"`
	IDUser       string  `json:"iduser"`
	IDTournament string  `json:"idtournament"`
	Type         string  `json:"type"`
	Value        float64 `json:"value"`
}
