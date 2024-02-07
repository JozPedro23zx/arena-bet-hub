package findparticipant

type ParcitipantInputDto struct {
	ID string `json:"id"`
}

type ParticipantOutputDto struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	NickName      string   `json:"nickname"`
	CountryOrigin string   `json:"countryorigin"`
	Tournmaents   []string `json:"tournaments"`
}
