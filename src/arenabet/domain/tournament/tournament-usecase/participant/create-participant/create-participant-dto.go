package createparticipant

type ParcitipantInputDto struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	NickName      string `json:"nickname"`
	CountryOrigin string `json:"countryorigin"`
}

type ParticipantOutputDto struct {
	ID       string `json:"id"`
	NickName string `json:"nickname"`
}
