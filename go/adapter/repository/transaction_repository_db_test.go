package repository

import (
	"os"
	"testing"

	"github.com/HRsniper/imersao-fullstack-fullcycle-5/adapter/repository/fixture"
	"github.com/HRsniper/imersao-fullstack-fullcycle-5/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestTransactionRepositoryDb_Insert(t *testing.T) {
	migrationsDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)
	repository := NewTransactionRepositoryDb(db)
	err := repository.Insert("1", "1", 12.1, entity.APPROVED, "")
	assert.Nil(t, err)
}
