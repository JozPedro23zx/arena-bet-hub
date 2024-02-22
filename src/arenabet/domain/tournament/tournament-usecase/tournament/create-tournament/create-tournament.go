package createtournament

import (
	"errors"

	Tournament "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-entities"
	repository "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-repositories"
	"github.com/JozPedro23zx/arena-bet-hub/infrastructure/broker"
)

type CreateTournament struct {
	Repository repository.TournamentRepository
	Producer   broker.ProducerInterface
	Topic      string
}

func NewCreateTournament(repository repository.TournamentRepository, producer broker.ProducerInterface, topic string) *CreateTournament {
	return &CreateTournament{Repository: repository, Producer: producer, Topic: topic}
}

func (ct *CreateTournament) Execute(input TournamentInputDto) (TournamentOutputDto, error) {
	tournament, err := ct.Repository.Find(input.ID)

	if err != nil {
		return TournamentOutputDto{}, err
	}

	if tournament != nil {
		output := TournamentOutputDto{
			ID:   tournament.ID,
			Name: tournament.Name,
		}
		err := errors.New("tournament already exist")

		return output, err
	}

	location := Tournament.Location{
		Street:  input.Street,
		City:    input.City,
		State:   input.State,
		Country: input.Country,
	}

	newTournament := Tournament.NewTournament(input.ID, input.Name, input.EventDate, location)

	err = ct.Repository.Insert(*newTournament)

	if err != nil {
		return TournamentOutputDto{}, err
	}
	output := TournamentOutputDto{
		ID:       newTournament.ID,
		Name:     newTournament.Name,
		Finished: newTournament.Finished,
	}

	err = ct.publish(output, []byte(newTournament.ID))
	if err != nil {
		return TournamentOutputDto{}, err
	}

	return output, nil
}

func (ct *CreateTournament) publish(output TournamentOutputDto, key []byte) error {
	err := ct.Producer.Publish(output, key, ct.Topic)
	if err != nil {
		return err
	}
	return nil
}
