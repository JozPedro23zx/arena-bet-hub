package tournamentkafka

import (
	"encoding/json"

	createtournament "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-usecase/tournament/create-tournament"
)

type CreateTournamentPresenter struct {
	ID       string `json:"id"`
	Finished bool   `json:"finished"`
}

func NewCreateTournamentPresenter() *CreateTournamentPresenter {
	return &CreateTournamentPresenter{}
}

func (t *CreateTournamentPresenter) Bind(result interface{}) error {
	t.ID = result.(createtournament.TournamentOutputDto).ID
	t.Finished = result.(createtournament.TournamentOutputDto).Finished
	return nil
}

func (t *CreateTournamentPresenter) Show() ([]byte, error) {
	j, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return j, nil
}
