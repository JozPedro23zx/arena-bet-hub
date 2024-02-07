package updatetournament

import (
	Tournament "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-entities"
	repository "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-repositories"
)

type UpdateTournament struct {
	Repository repository.TournamentRepository
}

func NewUpdateTournament(repository repository.TournamentRepository) *UpdateTournament {
	return &UpdateTournament{Repository: repository}
}

func (ut *UpdateTournament) Execute(input TournamentInputDto) (TournamentOutputDto, error) {
	tournament, err := ut.Repository.Find(input.ID)

	if err != nil {
		return TournamentOutputDto{}, err
	}

	newLocation := Tournament.Location{
		Street:  input.Street,
		City:    input.City,
		State:   input.State,
		Country: input.Country,
	}

	err = tournament.UpdateTournament(input.Name, input.EventDate, newLocation)

	if err != nil {
		return TournamentOutputDto{}, err
	}

	updatedTournament, err := ut.Repository.Update(*tournament)

	if err != nil {
		return TournamentOutputDto{}, err
	}

	output := TournamentOutputDto{
		ID:           updatedTournament.ID,
		Name:         updatedTournament.Name,
		EventDate:    updatedTournament.EventDate,
		Street:       updatedTournament.Location.Street,
		City:         updatedTournament.Location.City,
		State:        updatedTournament.Location.State,
		Country:      updatedTournament.Location.Country,
		Participants: updatedTournament.Participants(),
		Finished:     updatedTournament.Finished,
	}

	return output, nil
}

func (ut *UpdateTournament) AddParticipant(input ParticipantListDto) (TournamentOutputDto, error) {
	tournament, err := ut.Repository.Find(input.IDTournament)

	if err != nil {
		return TournamentOutputDto{}, err
	}

	if !input.Add {
		tournament.RemoveParticipant(input.ParticipantID)
	} else {
		participantExist := tournament.RegisterParticipant(input.ParticipantID)

		if participantExist != nil {
			return TournamentOutputDto{}, participantExist
		}
	}

	updatedTournament, err := ut.Repository.Update(*tournament)

	if err != nil {
		return TournamentOutputDto{}, err
	}

	output := TournamentOutputDto{
		ID:           updatedTournament.ID,
		Name:         updatedTournament.Name,
		EventDate:    updatedTournament.EventDate,
		Street:       updatedTournament.Location.Street,
		City:         updatedTournament.Location.City,
		State:        updatedTournament.Location.State,
		Country:      updatedTournament.Location.Country,
		Participants: updatedTournament.Participants(),
		Finished:     updatedTournament.Finished,
	}

	return output, nil
}
