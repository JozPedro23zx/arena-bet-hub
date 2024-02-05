package findresult

import (
	repository "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-repositories"
)

type FindResult struct {
	Repository repository.ResultRepository
}

func NewFindResult(repository repository.ResultRepository) *FindResult {
	return &FindResult{Repository: repository}
}

func (fr *FindResult) Execute(input ResultInputDto) (ResultOutputDto, error) {
	result, err := fr.Repository.Find(input.ID)

	if err != nil {
		return ResultOutputDto{}, err
	}

	output := ResultOutputDto{
		ID:          result.ID,
		TurnamentID: result.TournamentId,
		Open:        result.Open,
	}

	return output, nil
}
