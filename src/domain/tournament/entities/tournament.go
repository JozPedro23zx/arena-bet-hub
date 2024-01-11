package tournament

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
	ID             int
	Name           string
	EventDate      time.Time
	Location       Location
	IDParticipants []int
	Finished       bool
}

func NewTournament(ID int, name string, eventDate time.Time, location Location) *Tournament {
	return &Tournament{
		ID:             ID,
		Name:           name,
		EventDate:      eventDate,
		Location:       location,
		IDParticipants: []int{},
		Finished:       false,
	}
}

func (t *Tournament) UpdateTournament(name string, eventDate time.Time, location Location) error {
	if !t.Finished {
		t.Name = name
		t.EventDate = eventDate
		t.Location = location

		return nil
	}

	return errors.New("This tournament has finished")
}

func (t *Tournament) RegisterParticipant(participant Participant) error {
	if !t.Finished {
		t.IDParticipants = append(t.IDParticipants, participant.ID)
		return nil
	}
	return errors.New("This tournament has finished")
}

func (t *Tournament) RemoveParticipant(participantID int) error {
	if !t.Finished {
		for i, id := range t.IDParticipants {
			if id == participantID {
				t.IDParticipants = append(t.IDParticipants[:i], t.IDParticipants[i+1:]...)
				break
			}
		}
		return nil
	}
	return errors.New("This tournament has finished")
}

func (t *Tournament) FinishTournament() {
	t.Finished = true
}
