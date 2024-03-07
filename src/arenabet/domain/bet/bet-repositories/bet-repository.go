package betrepositories

import Bet "github.com/JozPedro23zx/arena-bet-hub/domain/bet/bet-entities"

type BetRepository interface {
	Insert(Bet.Bet) error
	Find(id string) (*Bet.Bet, error)
}
