package tournament_repositories

import Participant "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-entities"

type ParticipantRepository interface {
	Insert(participant Participant.Participant) error
	Find(id string) (*Participant.Participant, error)
	Update(participant Participant.Participant) (*Participant.Participant, error)
}
