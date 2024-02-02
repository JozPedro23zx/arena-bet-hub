package createparticipant

import (
	"errors"

	Participant "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-entities"
	repository "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-repositories"
)

type CreateParticipant struct {
	Repository repository.ParticipantRepository
}

func NewCreateParticipant(repository repository.ParticipantRepository) *CreateParticipant {
	return &CreateParticipant{Repository: repository}
}

func (cp *CreateParticipant) Execute(input ParcitipantInputDto) (ParticipantOutputDto, error) {
	participant, err := cp.Repository.Find(input.ID)

	if err != nil {
		return ParticipantOutputDto{}, err
	}

	if participant != nil {
		output := ParticipantOutputDto{
			ID:       participant.ID,
			NickName: participant.NickName,
		}
		err := errors.New("Participant already exist")

		return output, err
	}

	newParticipant := Participant.NewParticipant(input.ID, input.Name, input.NickName, input.CountryOrigin)

	err = cp.Repository.Insert(*newParticipant)

	if err != nil {
		return ParticipantOutputDto{}, err
	}

	output := ParticipantOutputDto{
		ID:       input.ID,
		NickName: input.NickName,
	}

	return output, err
}
