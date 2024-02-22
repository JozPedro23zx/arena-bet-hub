package updateresult

import (
	repository "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-repositories"
	"github.com/JozPedro23zx/arena-bet-hub/infrastructure/broker"
)

type CloseOrOpenResult struct {
	Repository repository.ResultRepository
	Producer   broker.ProducerInterface
	Topic      string
}

func NewCloseOrOpenResult(repository repository.ResultRepository, producer broker.ProducerInterface, topic string) *CloseOrOpenResult {
	return &CloseOrOpenResult{Repository: repository, Producer: producer, Topic: topic}
}

func (cr *CloseOrOpenResult) Execute(input OpenInputDto) (ResultOutputDto, error) {
	result, err := cr.Repository.Find(input.ID)

	if err != nil {
		return ResultOutputDto{}, err
	}

	if !input.Open {
		result.CloseResult()
	}

	updatedResult, err := cr.Repository.Update(*result)

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

	err = cr.publish(output, []byte(input.ID))
	if err != nil {
		return ResultOutputDto{}, err
	}

	return output, nil
}

func (cr *CloseOrOpenResult) publish(result ResultOutputDto, key []byte) error {
	err := cr.Producer.Publish(result, key, cr.Topic)
	if err != nil {
		return err
	}
	return nil
}
