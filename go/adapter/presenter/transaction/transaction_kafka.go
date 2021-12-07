package transaction

import (
	"encoding/json"

	"github.com/HRsniper/imersao-fullstack-fullcycle-5/usecase/process_transaction"
)

type KafkaPresenter struct {
	ID           string `json:"id"`
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}

// constructor
func NewTransactionKafkaPresenter() *KafkaPresenter {
	return &KafkaPresenter{}
}

// public method
func (t *KafkaPresenter) Bind(input interface{}) error {
	t.ID = input.(process_transaction.TransactionDtoOutput).ID
	t.Status = input.(process_transaction.TransactionDtoOutput).Status
	t.ErrorMessage = input.(process_transaction.TransactionDtoOutput).ErrorMessage

	return nil
}

// public method
func (t *KafkaPresenter) Show() ([]byte, error) {
	j, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	return j, nil
}