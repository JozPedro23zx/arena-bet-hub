package createresult

import (
	"errors"

	repository "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-repositories"

	Tournament "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-entities"
)

type CreateResult struct {
	Repository repository.ResultRepository
}

func NewCreateResult(repository repository.ResultRepository) *CreateResult {
	return &CreateResult{Repository: repository}
}

func (cr *CreateResult) Execute(input ResultInputDto) (ResultOutputDto, error) {
	result, err := cr.Repository.Find(input.ID)

	if err != nil {
		return ResultOutputDto{}, err
	}

	if result != nil {
		output := ResultOutputDto{
			ID:          result.ID,
			TurnamentID: result.TournamentId,
			Open:        result.Open,
		}

		err := errors.New("Result alredy existe")

		return output, err
	}

	newResult := Tournament.NewResult(input.ID, input.TurnamentID)

	err = cr.Repository.Insert(*newResult)

	if err != nil {
		return ResultOutputDto{}, err
	}

	output := ResultOutputDto{
		ID:          newResult.ID,
		TurnamentID: newResult.TournamentId,
		Open:        newResult.Open,
	}

	return output, err
}
