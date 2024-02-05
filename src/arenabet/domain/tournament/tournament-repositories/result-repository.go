package tournament_repositories

import Result "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-entities"

type ResultRepository interface {
	Insert(result Result.Result) error
	Find(id string) (*Result.Result, error)
	Update(result Result.Result) (*Result.Result, error)
}
