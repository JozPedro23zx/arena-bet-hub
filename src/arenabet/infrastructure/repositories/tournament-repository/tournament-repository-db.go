package tournamentrepository

import (
	"database/sql"

	Tournament "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-entities"
)

type TournamentRepositoryDB struct {
	db *sql.DB
}

func NewTournamentRepositoryDB(db *sql.DB) *TournamentRepositoryDB {
	return &TournamentRepositoryDB{db: db}
}

func (t *TournamentRepositoryDB) Insert(tournament Tournament.Tournament) error {
	stmt, err := t.db.Prepare(`insert into tournaments (id, name, event_date, street, city, state, country, finished ) values(?, ?, ?, ?, ?, ?, ?, ?)`)

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
