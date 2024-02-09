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

	var tournamentID string
	var tournamentName string
	var tournamentDate string
	var street string
	var state string
	var city string
	var country string
	var tournamentFinished string

	err := row.Scan(
		&tournamentID,
		&tournamentName,
		&tournamentDate,
		&street,
		&state,
		&city,
		&country,
		&tournamentFinished,
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
		Street:  street,
		State:   state,
		City:    city,
		Country: country,
	}
	tournament := Tournament.NewTournament(tournamentID, tournamentName, eventDate, location)

	return tournament, nil
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
