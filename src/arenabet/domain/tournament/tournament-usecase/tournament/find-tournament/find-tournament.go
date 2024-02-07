package findtournament

import (
	repository "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-repositories"
)

type FindTournament struct {
	Repository repository.TournamentRepository
}

func NewFindTournament(repository repository.TournamentRepository) *FindTournament {
	return &FindTournament{Repository: repository}
}

func (ft *FindTournament) Execute(input TournamentInputDto) (TournamentOutputDto, error) {
	tournament, err := ft.Repository.Find(input.ID)

	if err != nil {
		return TournamentOutputDto{}, err
	}

	output := TournamentOutputDto{
		ID:           tournament.ID,
		Name:         tournament.Name,
		EventDate:    tournament.EventDate,
		Street:       tournament.Location.Street,
		City:         tournament.Location.City,
		State:        tournament.Location.State,
		Country:      tournament.Location.Country,
		Participants: tournament.Participants(),
		Finished:     tournament.Finished,
	}

	return output, nil
}
