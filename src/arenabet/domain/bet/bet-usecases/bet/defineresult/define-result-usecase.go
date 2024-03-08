package defineresult

import (
	"errors"

	betentities "github.com/JozPedro23zx/arena-bet-hub/domain/bet/bet-entities"
	betrepositories "github.com/JozPedro23zx/arena-bet-hub/domain/bet/bet-repositories"
)

type DefineResult struct {
	Repository betrepositories.BetResultRepository
}

func NewDefineResult(repository betrepositories.BetResultRepository) *DefineResult {
	return &DefineResult{Repository: repository}
}

func (dr *DefineResult) Execute(input BetInputDto) error {
	result, err := dr.Repository.Find(input.ID)

	if err != nil {
		return err
	}

	if result != nil {
		return errors.New("this result has been defined")
	}

	resultDefined := betentities.NewBetResult(input.ID, input.IDUser, input.IDBet, input.Payment, input.Victory, input.Status)

	err = dr.Repository.Insert(*resultDefined)

	if err != nil {
		return err
	}

	return nil
}
