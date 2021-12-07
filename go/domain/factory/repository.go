package factory

import "github.com/HRsniper/imersao-fullstack-fullcycle-5/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
