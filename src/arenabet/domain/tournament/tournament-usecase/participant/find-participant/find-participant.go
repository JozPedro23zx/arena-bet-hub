package findparticipant

import repository "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-repositories"

type FindParticipant struct {
	Repository repository.ParticipantRepository
}

func NewFindParticipant(repository repository.ParticipantRepository) *FindParticipant {
	return &FindParticipant{Repository: repository}
}

func (fp *FindParticipant) Execute(input ParcitipantInputDto) (ParticipantOutputDto, error) {
	participant, err := fp.Repository.Find(input.ID)

	if err != nil {
		return ParticipantOutputDto{}, err
	}

	output := ParticipantOutputDto{
		ID:            participant.ID,
		Name:          participant.Name,
		NickName:      participant.NickName,
		CountryOrigin: participant.CountryOrigin,
		Tournmaents:   participant.Tournamnets(),
	}

	return output, nil
}
