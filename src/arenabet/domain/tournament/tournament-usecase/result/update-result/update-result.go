package updateresult

import (
	"errors"

	repository "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-repositories"
)

type UpdateResult struct {
	Repository repository.ResultRepository
}

func NewUpdateResult(repository repository.ResultRepository) *UpdateResult {
	return &UpdateResult{Repository: repository}
}

func (ur *UpdateResult) CloseOrOpenResult(input OpenInputDto) (ResultOutputDto, error) {
	result, err := ur.Repository.Find(input.ID)

	if err != nil {
		return ResultOutputDto{}, err
	}

	if !input.Open {
		result.CloseResult()
	}

	updatedResult, err := ur.Repository.Update(*result)

	if err != nil {
		return ResultOutputDto{}, err
	}

	output := ResultOutputDto{
		ID:           updatedResult.ID,
		TurnamentID:  updatedResult.TournamentId,
		Ranking:      updatedResult.Ranking,
		Open:         updatedResult.Open,
		DateFinished: updatedResult.DateFinished,
	}

	return output, nil
}

func (ur *UpdateResult) AddParticipant(input RankingInputDto) error {
	result, err := ur.Repository.Find(input.ID)

	if err != nil {
		return err
	}

	_, err = result.DefineRanking(input.ParticipantID)

	if err != nil {
		return err
	}

	return nil
}

func (ur *UpdateResult) UpdateRanking(input RankingInputDto) (ResultOutputDto, error) {
	result, err := ur.Repository.Find(input.ID)

	if err != nil {
		return ResultOutputDto{}, err
	}

	if !result.Open {
		output := ResultOutputDto{
			ID:           result.ID,
			TurnamentID:  result.TournamentId,
			Ranking:      result.Ranking,
			Open:         result.Open,
			DateFinished: result.DateFinished,
		}
		err := errors.New("result has been closed")

		return output, err
	}

	err = result.UpdateRanking(input.ParticipantID, input.Score)

	if err != nil {
		return ResultOutputDto{}, err
	}

	updatedResult, err := ur.Repository.Update(*result)

	if err != nil {
		return ResultOutputDto{}, err
	}

	output := ResultOutputDto{
		ID:           updatedResult.ID,
		TurnamentID:  updatedResult.TournamentId,
		Ranking:      updatedResult.Ranking,
		Open:         updatedResult.Open,
		DateFinished: updatedResult.DateFinished,
	}

	return output, err
}
