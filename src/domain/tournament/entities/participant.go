package tournament

type Participant struct {
	ID            int
	Name          string
	NickName      string
	CountryOrigin string
	Tournaments   []int
}

func NewParticipant(id int, name string, nickName string, countryOrigin string) *Participant {
	return &Participant{
		ID:            id,
		Name:          name,
		NickName:      nickName,
		CountryOrigin: countryOrigin,
		Tournaments:   []int{},
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
