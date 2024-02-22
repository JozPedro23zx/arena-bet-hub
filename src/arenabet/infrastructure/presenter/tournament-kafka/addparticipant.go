package tournamentkafka

import (
	"encoding/json"

	updateTournament "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-usecase/tournament/update-tournament"
)

type AddParticipantPresenter struct {
	TournamentID  string `json:"tournamentid"`
	ParticipantID string `json:"participantid"`
}

func NewAddParticipantPresenter() *AddParticipantPresenter {
	return &AddParticipantPresenter{}
}

func (t *AddParticipantPresenter) Bind(result interface{}) error {
	t.TournamentID = result.(updateTournament.ParticipantDto).IDTournament
	t.ParticipantID = result.(updateTournament.ParticipantDto).ParticipantID
	return nil
}

func (t *AddParticipantPresenter) Show() ([]byte, error) {
	j, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return j, nil
}
