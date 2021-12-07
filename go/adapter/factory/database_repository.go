package factory

import (
	"database/sql"

	repo "github.com/HRsniper/imersao-fullstack-fullcycle-5/adapter/repository"
	"github.com/HRsniper/imersao-fullstack-fullcycle-5/domain/repository"
)

type RepositoryDatabaseFactory struct {
	DB *sql.DB
}

// constructor
func NewRepositoryDatabaseFactory(db *sql.DB) *RepositoryDatabaseFactory {
	return &RepositoryDatabaseFactory{DB: db}
}

// public method
func (r RepositoryDatabaseFactory) CreateTransactionRepository() repository.TransactionRepository {
	return repo.NewTransactionRepositoryDb(r.DB)
}
