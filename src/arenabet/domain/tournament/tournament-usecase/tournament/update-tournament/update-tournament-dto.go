package updatetournament

import "time"

type TournamentInputDto struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	EventDate time.Time `json:"eventDate"`
	Street    string    `json:"street"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	Country   string    `json:"country"`
}

type ParticipantListDto struct {
	IDTournament  string `json:"idtournament"`
	ParticipantID string `json:"participantid"`
}

type TournamentOutputDto struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	EventDate    time.Time `json:"eventDate"`
	Street       string    `json:"street"`
	City         string    `json:"city"`
	State        string    `json:"state"`
	Country      string    `json:"country"`
	Participants []string  `json:"participants"`
	Finished     bool      `json:"finished"`
}
