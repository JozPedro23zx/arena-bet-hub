package betrepositories

import Bet "github.com/JozPedro23zx/arena-bet-hub/domain/bet/bet-entities"

type BetResultRepository interface {
	Insert(Bet.BetResult) error
	Find(id string) (*Bet.BetResult, error)
}
