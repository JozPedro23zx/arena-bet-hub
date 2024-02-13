package participantrepository

import (
	"database/sql"
	"errors"

	Tournament "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-entities"
)

type ParticipantRepositoryDB struct {
	db *sql.DB
}

func NewParticipantRepositoryDB(db *sql.DB) *ParticipantRepositoryDB {
	return &ParticipantRepositoryDB{db: db}
}

func (p *ParticipantRepositoryDB) Insert(participant Tournament.Participant) error {
	stmt, err := p.db.Prepare(`INSERT into participants (id, name, nick_name, country_origin) values(?, ?, ?, ?)`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		participant.ID,
		participant.Name,
		participant.NickName,
		participant.CountryOrigin,
	)

	if err != nil {
		return err
	}

	return nil
}

func (p *ParticipantRepositoryDB) Find(id string) (*Tournament.Participant, error) {
	query := `SELECT id, name, nick_name, country_origin FROM participants WHERE id = ?`
	row := p.db.QueryRow(query, id)

	participant := Tournament.Participant{}

	err := row.Scan(
		&participant.ID,
		&participant.Name,
		&participant.NickName,
		&participant.CountryOrigin,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("participant not found")
		}
		return nil, err
	}

	participantFound := Tournament.NewParticipant(participant.ID, participant.Name, participant.NickName, participant.CountryOrigin)

	return participantFound, nil
}

func (p *ParticipantRepositoryDB) Update(participant Tournament.Participant) (*Tournament.Participant, error) {
	stmt, err := p.db.Prepare(`UPDATE participants SET name=?, nick_name=?, country_origin=? WHERE id=?`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		participant.Name,
		participant.NickName,
		participant.CountryOrigin,
		participant.ID,
	)

	if err != nil {
		return nil, err
	}

	updatedParticipant, err := p.Find(participant.ID)

	if err != nil {
		return nil, err
	}

	return updatedParticipant, nil

}
