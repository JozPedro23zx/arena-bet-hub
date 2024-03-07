package defineresult

import (
	"testing"

	betentities "github.com/JozPedro23zx/arena-bet-hub/domain/bet/bet-entities"
	mock_betrepositories "github.com/JozPedro23zx/arena-bet-hub/domain/bet/bet-repositories/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDefineResut(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositorymock := mock_betrepositories.NewMockBetResultRepository(ctrl)

	input := BetInputDto{
		ID:      "result123",
		IDUser:  "user123",
		IDBet:   "bet123",
		Payment: 4000,
		Victory: true,
		Status:  false,
	}

	betResult := betentities.NewBetResult(input.ID, input.IDUser, input.IDBet, input.Payment, input.Victory, input.Status)

	repositorymock.EXPECT().Find(input.ID).Return(nil, nil)
	repositorymock.EXPECT().Insert(*betResult).Return(nil)

	defineResultUsecase := NewDefineResult(repositorymock)
	err := defineResultUsecase.Execute(input)

	assert.Nil(t, err)
}
