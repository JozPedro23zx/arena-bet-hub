package resultrepository

import (
	"database/sql"
	"errors"
	"time"

	Tournament "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-entities"
)

type ResultRepositoryDb struct {
	db *sql.DB
}

func NewResultRepositoryDB(db *sql.DB) *ResultRepositoryDb {
	return &ResultRepositoryDb{db: db}
}

func (r *ResultRepositoryDb) Insert(result Tournament.Result) error {
	stmt, err := r.db.Prepare(`INSERT into results (id, tournament_id, open, date_finished) values(?, ?, ?, ?)`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		result.ID,
		result.TournamentId,
		result.Open,
		result.DateFinished,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *ResultRepositoryDb) Find(id string) (*Tournament.Result, error) {
	query := `SELECT id, tournament_id, open, date_finished FROM results WHERE id = ?`
	row := r.db.QueryRow(query, id)

	result := Tournament.Result{}
	var resultDateFinished string

	err := row.Scan(
		&result.ID,
		&result.TournamentId,
		&result.Open,
		&resultDateFinished,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("result not found")
		}
		return nil, err
	}

	dateFinished, err := time.Parse("2006-01-02 15:04:05.999999999-07:00", resultDateFinished)
	if err != nil {
		return nil, err
	}

	resultFound := Tournament.NewResult(result.ID, result.TournamentId)

	if !result.Open {
		resultFound.CloseResult()
		resultFound.DateFinished = dateFinished
	}

	return resultFound, nil
}

func (r *ResultRepositoryDb) Update(result Tournament.Result) (*Tournament.Result, error) {
	stmt, err := r.db.Prepare(`UPDATE results SET tournament_id = ?, open = ?, date_finished = ? WHERE id = ?`)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(
		&result.TournamentId,
		&result.Open,
		&result.DateFinished,
		&result.ID,
	)

	if err != nil {
		return nil, err
	}

	updatedResult, err := r.Find(result.ID)

	if err != nil {
		return nil, err
	}

	return updatedResult, nil
}
