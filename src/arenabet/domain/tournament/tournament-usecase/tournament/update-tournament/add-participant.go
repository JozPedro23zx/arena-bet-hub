package updatetournament

import (
	repository "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-repositories"
	"github.com/JozPedro23zx/arena-bet-hub/infrastructure/broker"
)

type AddParticipant struct {
	Repository repository.TournamentRepository
	Producer   broker.ProducerInterface
	Topic      string
}

func NewAddParticipant(repository repository.TournamentRepository, poroducer broker.ProducerInterface, topic string) *AddParticipant {
	return &AddParticipant{Repository: repository, Producer: poroducer, Topic: topic}
}

func (ap *AddParticipant) Execute(input ParticipantDto) (TournamentOutputDto, error) {
	tournament, err := ap.Repository.Find(input.IDTournament)

	if err != nil {
		return TournamentOutputDto{}, err
	}

	if !input.Add {
		tournament.RemoveParticipant(input.ParticipantID)
	} else {
		participantExist := tournament.RegisterParticipant(input.ParticipantID)

		if participantExist != nil {
			return TournamentOutputDto{}, participantExist
		}
	}

	updatedTournament, err := ap.Repository.Update(*tournament)

	if err != nil {
		return TournamentOutputDto{}, err
	}

	output := TournamentOutputDto{
		ID:           updatedTournament.ID,
		Name:         updatedTournament.Name,
		EventDate:    updatedTournament.EventDate,
		Street:       updatedTournament.Location.Street,
		City:         updatedTournament.Location.City,
		State:        updatedTournament.Location.State,
		Country:      updatedTournament.Location.Country,
		Participants: updatedTournament.Participants(),
		Finished:     updatedTournament.Finished,
	}

	err = ap.publish(output, []byte(updatedTournament.ID))

	if err != nil {
		return TournamentOutputDto{}, err
	}

	return output, nil
}

func (ap *AddParticipant) publish(output TournamentOutputDto, key []byte) error {
	err := ap.Producer.Publish(output, key, ap.Topic)
	if err != nil {
		return err
	}
	return nil
}
