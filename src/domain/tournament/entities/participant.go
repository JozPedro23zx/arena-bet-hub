package tournament

type Participant struct {
	ID            int
	Name          string
	NickName      string
	CountryOrigin string
}

func NewParticipant(id int, name string, nickName string, countryOrigin string) *Participant {
	return &Participant{
		ID:            id,
		Name:          name,
		NickName:      nickName,
		CountryOrigin: countryOrigin,
	}
}

func (p *Participant) UpdateParticipant(name string, nickName string, countryOrigin string) {
	p.Name = name
	p.NickName = nickName
	p.CountryOrigin = countryOrigin
}
