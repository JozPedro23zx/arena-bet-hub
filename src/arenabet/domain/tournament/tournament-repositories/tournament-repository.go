package tournament_repositories

import Tournament "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-entities"

type TournamentRepository interface {
	Insert(tournament Tournament.Tournament) error
	Find(id string) (*Tournament.Tournament, error)
	Update(tournament Tournament.Tournament) (*Tournament.Tournament, error)
}
