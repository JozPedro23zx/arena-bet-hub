package tournament_entities

import "errors"

type Participant struct {
	ID            string
	Name          string
	NickName      string
	CountryOrigin string
	tournaments   []string
}

func NewParticipant(id string, name string, nickName string, countryOrigin string) *Participant {
	return &Participant{
		ID:            id,
		Name:          name,
		NickName:      nickName,
		CountryOrigin: countryOrigin,
		tournaments:   []string{},
	}
}

func (p *Participant) UpdateParticipant(name string, nickName string, countryOrigin string) {
	p.Name = name
	p.NickName = nickName
	p.CountryOrigin = countryOrigin
}

func (p *Participant) RegisterTournamentParticipation(tournamentId string) error {
	for _, existingId := range p.tournaments {
		if existingId == tournamentId {
			return errors.New("tournament already exist")
		}
	}
	p.tournaments = append(p.tournaments, tournamentId)
	return nil
}

func (p *Participant) Tournamnets() []string {
	return p.tournaments
}
