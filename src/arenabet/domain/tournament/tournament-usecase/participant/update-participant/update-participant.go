package findparticipant

import (
	repository "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-repositories"
)

type UpdateParticipant struct {
	Repository repository.ParticipantRepository
}

func NewUpdateParticipant(repository repository.ParticipantRepository) *UpdateParticipant {
	return &UpdateParticipant{Repository: repository}
}

func (up UpdateParticipant) Execute(input ParcitipantInputDto) (ParticipantOutputDto, error) {
	participant, err := up.Repository.Find(input.ID)

	if err != nil {
		return ParticipantOutputDto{}, err
	}

	participant.UpdateParticipant(input.Name, input.NickName, input.CountryOrigin)

	participantUpdated, err := up.Repository.Update(*participant)

	if err != nil {
		return ParticipantOutputDto{}, err
	}

	output := ParticipantOutputDto{
		ID:            participantUpdated.ID,
		Name:          participant.Name,
		NickName:      participant.NickName,
		CountryOrigin: participant.CountryOrigin,
	}

	return output, nil
}
