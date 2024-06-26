package tournamentrepository

import (
	"database/sql"
	"errors"
	"time"

	Tournament "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-entities"
)

type TournamentRepositoryDB struct {
	db *sql.DB
}

func NewTournamentRepositoryDB(db *sql.DB) *TournamentRepositoryDB {
	return &TournamentRepositoryDB{db: db}
}

func (t *TournamentRepositoryDB) Insert(tournament Tournament.Tournament) error {
	stmt, err := t.db.Prepare(`INSERT into tournaments (id, name, event_date, street, city, state, country, finished ) values(?, ?, ?, ?, ?, ?, ?, ?)`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		tournament.ID,
		tournament.Name,
		tournament.EventDate,
		tournament.Location.Street,
		tournament.Location.City,
		tournament.Location.State,
		tournament.Location.Country,
		tournament.Finished,
	)

	if err != nil {
		return err
	}

	return nil
}

func (t *TournamentRepositoryDB) Find(id string) (*Tournament.Tournament, error) {
	query := `SELECT id, name, event_date, street, state, city, country, finished FROM tournaments WHERE id = ?`
	row := t.db.QueryRow(query, id)

	tournament := Tournament.Tournament{}
	var tournamentDate string

	err := row.Scan(
		&tournament.ID,
		&tournament.Name,
		&tournamentDate,
		&tournament.Location.Street,
		&tournament.Location.State,
		&tournament.Location.City,
		&tournament.Location.Country,
		&tournament.Finished,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("tournament not found")
		}
		return nil, err
	}

	eventDate, err := time.Parse("2006-01-02 15:04:05.999999999-07:00", tournamentDate)
	if err != nil {
		return nil, err
	}

	location := Tournament.Location{
		Street:  tournament.Location.Street,
		State:   tournament.Location.State,
		City:    tournament.Location.City,
		Country: tournament.Location.Country,
	}
	tournamentFound := Tournament.NewTournament(tournament.ID, tournament.Name, eventDate, location)

	return tournamentFound, nil
}

func (t *TournamentRepositoryDB) Update(tournament Tournament.Tournament) (*Tournament.Tournament, error) {
	stmt, err := t.db.Prepare(`UPDATE tournaments SET name=?, event_date=?, street=?, city=?, state=?, country=?, finished=? WHERE id=?`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		tournament.Name,
		tournament.EventDate,
		tournament.Location.Street,
		tournament.Location.City,
		tournament.Location.State,
		tournament.Location.Country,
		tournament.Finished,
		tournament.ID,
	)
	if err != nil {
		return nil, err
	}

	updatedTournament, err := t.Find(tournament.ID)
	if err != nil {
		return nil, err
	}

	return updatedTournament, nil
}
