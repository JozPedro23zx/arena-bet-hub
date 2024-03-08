package createbet

import (
	"errors"

	betentities "github.com/JozPedro23zx/arena-bet-hub/domain/bet/bet-entities"
	betrepositories "github.com/JozPedro23zx/arena-bet-hub/domain/bet/bet-repositories"
)

type CreateBet struct {
	Repository betrepositories.BetRepository
}

func NewCreateBet(repository betrepositories.BetRepository) *CreateBet {
	return &CreateBet{Repository: repository}
}

func (cb *CreateBet) Execute(input BetInputDto) error {
	bet, err := cb.Repository.Find(input.ID)

	if err != nil {
		return err
	}

	if bet != nil {
		return errors.New("bet already exist")
	}

	newBet := betentities.NewBet(input.ID, input.IDUser, input.IDTournament, input.Value)

	err = cb.Repository.Insert(*newBet)

	if err != nil {
		return err
	}

	return nil
}
