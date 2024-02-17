package tournamentkafka

import (
	"encoding/json"

	updateResult "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-usecase/result/update-result"
)

type CloseResultPresenter struct {
	ID   string `json:"tournamentid"`
	Open bool   `json:"participantid"`
}

func NewCloseResultPresenter() *CloseResultPresenter {
	return &CloseResultPresenter{}
}

func (t *CloseResultPresenter) Bind(result interface{}) error {
	t.ID = result.(updateResult.ResultOutputDto).ID
	t.Open = result.(updateResult.ResultOutputDto).Open
	return nil
}

func (t *CloseResultPresenter) Show() ([]byte, error) {
	j, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return j, nil
}
