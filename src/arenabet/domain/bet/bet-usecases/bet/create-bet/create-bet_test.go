package createbet

import (
	"testing"

	betentities "github.com/JozPedro23zx/arena-bet-hub/domain/bet/bet-entities"
	mock_betrepositories "github.com/JozPedro23zx/arena-bet-hub/domain/bet/bet-repositories/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateBet(t *testing.T) {
	cntrl := gomock.NewController(t)
	defer cntrl.Finish()
	repositorymock := mock_betrepositories.NewMockBetRepository(cntrl)

	input := BetInputDto{
		ID:           "bet123",
		IDUser:       "user123",
		IDTournament: "tournament123",
		Value:        200.0,
	}

	bet := betentities.NewBet(input.ID, input.IDUser, input.IDTournament, input.Value)

	repositorymock.EXPECT().Find(input.ID).Return(nil, nil)
	repositorymock.EXPECT().Insert(*bet).Return(nil)

	createBetUsecase := NewCreateBet(repositorymock)
	err := createBetUsecase.Execute(input)

	assert.Nil(t, err)
}
