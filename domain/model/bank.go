package model

import "time"
import "github.com/satori/go.uuid"

type Bank struct {
	ID        string    `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (receiver *Bank) validator() error {
	return nil
}

func NewBank(code string, name string) (*Bank, error) {
	bank := Bank{Code: code, Name: name}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()

	error := bank.validator()

	if error != nil {
		return nil, error
	}
	return &bank, nil
}
