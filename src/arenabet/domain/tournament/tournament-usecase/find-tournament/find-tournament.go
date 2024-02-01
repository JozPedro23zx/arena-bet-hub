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

func (ft *FindTournament) Execute(input TournamentInputDto) (TournamentOutupttDto, error) {
	tournament, err := ft.Repository.Find(input.ID)

	if err != nil {
		return TournamentOutupttDto{}, err
	}

	output := TournamentOutupttDto{
		ID:        tournament.ID,
		Name:      tournament.Name,
		EventDate: tournament.EventDate,
		Street:    tournament.Location.Street,
		City:      tournament.Location.City,
		State:     tournament.Location.State,
		Country:   tournament.Location.Country,
		Finished:  tournament.Finished,
	}

	return output, nil
}
