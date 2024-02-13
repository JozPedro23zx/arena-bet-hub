package resultrepository

import (
	"os"
	"testing"

	Tournament "github.com/JozPedro23zx/arena-bet-hub/domain/tournament/tournament-entities"
	"github.com/JozPedro23zx/arena-bet-hub/infrastructure/repositories/fixture"
	"github.com/stretchr/testify/assert"
)

func TestResultInsertDb(t *testing.T) {
	migrationsDir := os.DirFS("../fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)
	repository := NewResultRepositoryDB(db)

	result := Tournament.NewResult("result123", "tournament123")

	err := repository.Insert(*result)

	assert.Nil(t, err)
}

func TestResultFindDb(t *testing.T) {
	migrationsDir := os.DirFS("../fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)
	repository := NewResultRepositoryDB(db)

	result := Tournament.NewResult("result123", "tournament123")
	repository.Insert(*result)

	resultFound, err := repository.Find(result.ID)

	assert.Nil(t, err)
	assert.Equal(t, result, resultFound)
}

func TestResultUpdateDb(t *testing.T) {
	migrationsDir := os.DirFS("../fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)
	repository := NewResultRepositoryDB(db)

	result := Tournament.NewResult("result123", "tournament123")
	repository.Insert(*result)

	newResult := Tournament.NewResult(result.ID, result.TournamentId)
	newResult.CloseResult()

	resultUpdated, err := repository.Update(*newResult)

	assert.Nil(t, err)
	assert.Equal(t, newResult, resultUpdated)
	assert.NotEqual(t, result, resultUpdated)
}
