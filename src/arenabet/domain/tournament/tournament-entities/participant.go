package tournament_entities

type Participant struct {
	ID            string
	Name          string
	NickName      string
	CountryOrigin string
	Tournaments   []string
}

func NewParticipant(id string, name string, nickName string, countryOrigin string) *Participant {
	return &Participant{
		ID:            id,
		Name:          name,
		NickName:      nickName,
		CountryOrigin: countryOrigin,
		Tournaments:   []string{},
	}
}

func (p *Participant) UpdateParticipant(name string, nickName string, countryOrigin string) {
	p.Name = name
	p.NickName = nickName
	p.CountryOrigin = countryOrigin
}

func (p *Participant) RegisterTournamentParticipation(tournament Tournament) {
	p.Tournaments = append(p.Tournaments, tournament.ID)
}
