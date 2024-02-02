package updatetournament

import (
	Tournament "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-entities"
	repository "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-repositories"
)

type UpdateTournament struct {
	Repository repository.TournamentRepository
}

func NewUpdateTournament(repository repository.TournamentRepository) *UpdateTournament {
	return &UpdateTournament{Repository: repository}
}

func (ut *UpdateTournament) Execute(input TournamentInputDto) (TournamentOutputDto, error) {
	tournament, err := ut.Repository.Find(input.ID)

	if err != nil {
		return TournamentOutputDto{}, err
	}

	newLocation := Tournament.Location{
		Street:  input.Street,
		City:    input.City,
		State:   input.State,
		Country: input.Country,
	}

	tournament.UpdateTournament(input.Name, input.EventDate, newLocation)

	updatedTournament, err := ut.Repository.Update(*tournament)

	if err != nil {
		return TournamentOutputDto{}, err
	}

	output := TournamentOutputDto{
		ID:        updatedTournament.ID,
		Name:      updatedTournament.Name,
		EventDate: updatedTournament.EventDate,
		Street:    updatedTournament.Location.Street,
		City:      updatedTournament.Location.City,
		State:     updatedTournament.Location.State,
		Country:   updatedTournament.Location.Country,
		Finished:  updatedTournament.Finished,
	}

	return output, nil
}
