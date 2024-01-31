package createtournament

import (
	"errors"

	Tournament "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-entities"
	repository "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-repositories"
)

type CreateTournament struct {
	Repository repository.TournamentRepository
}

func NewCreateTournament(repository repository.TournamentRepository) *CreateTournament {
	return &CreateTournament{Repository: repository}
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
		err := errors.New("Tournament already exist")

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
		ID:   input.ID,
		Name: input.Name,
	}

	return output, nil
}
