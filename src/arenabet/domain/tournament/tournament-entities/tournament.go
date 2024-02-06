package tournament_entities

import (
	"errors"
	"time"
)

type Location struct {
	Street  string
	City    string
	State   string
	Country string
}

type Tournament struct {
	ID           string
	Name         string
	EventDate    time.Time
	Location     Location
	participants []string
	Finished     bool
}

func NewTournament(ID string, name string, eventDate time.Time, location Location) *Tournament {
	return &Tournament{
		ID:           ID,
		Name:         name,
		EventDate:    eventDate,
		Location:     location,
		participants: []string{},
		Finished:     false,
	}
}

func (t *Tournament) Participants() []string {
	return t.participants
}

func (t *Tournament) UpdateTournament(name string, eventDate time.Time, location Location) error {
	if !t.Finished {
		t.Name = name
		t.EventDate = eventDate
		t.Location = location

		return nil
	}

	return errors.New("this tournament has finished")
}

func (t *Tournament) RegisterParticipant(participantID string) error {
	if !t.Finished {
		for _, existingId := range t.participants {
			if existingId == participantID {
				return errors.New("participant already registered")
			}
		}
		t.participants = append(t.participants, participantID)
		return nil
	}
	return errors.New("this tournament has finished")
}

func (t *Tournament) RemoveParticipant(participantID string) error {
	if !t.Finished {
		for i, id := range t.participants {
			if id == participantID {
				t.participants = append(t.participants[:i], t.participants[i+1:]...)
				break
			}
		}
		return nil
	}
	return errors.New("this tournament has finished")
}

func (t *Tournament) FinishTournament() {
	t.Finished = true
}
