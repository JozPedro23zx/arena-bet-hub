package updateresult

import (
	"errors"

	Tournament "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-entities"
	repository "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-repositories"
)

type UpdateResult struct {
	Repository repository.ResultRepository
}

func NewUpdateResult(repository repository.ResultRepository) *UpdateResult {
	return &UpdateResult{Repository: repository}
}

func (ur *UpdateResult) AddParticipant(input RankingInputDto) error {
	result, err := ur.Repository.Find(input.ResultID)

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
	result, err := ur.Repository.Find(input.ResultID)

	if err != nil {
		return ResultOutputDto{}, err
	}

	if !result.Open {

		rankingOutput := getRanking(*result)
		output := ResultOutputDto{
			ID:           result.ID,
			TurnamentID:  result.TournamentId,
			Ranking:      rankingOutput,
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

	rankingOutput := getRanking(*updatedResult)
	output := ResultOutputDto{
		ID:           updatedResult.ID,
		TurnamentID:  updatedResult.TournamentId,
		Ranking:      rankingOutput,
		Open:         updatedResult.Open,
		DateFinished: updatedResult.DateFinished,
	}

	return output, err
}

func getRanking(result Tournament.Result) []RankingOutputDto {
	rankingOutput := []RankingOutputDto{}
	for _, rank := range result.Ranking() {
		output := RankingOutputDto{
			ParticipantID: rank.ParticipantId,
			Score:         rank.Score,
		}
		rankingOutput = append(rankingOutput, output)
	}
	return rankingOutput
}
