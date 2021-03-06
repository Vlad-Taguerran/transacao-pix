package model

import "time"
import "github.com/satori/go.uuid"
import "github.com/asaskevich/govalidator"

type Bank struct {
	Base     `valid:"requied"`
	Code     string     `json:"code" valid:"notnull"`
	Name     string     `json:"name" valid:"notnull"`
	Accounts []*Account `valid:"-"`
}

func (receiver *Bank) validator() error {
	_, err := govalidator.ValidateStruct(receiver)

	if err != nil {
		return err
	}
	return nil
}

func NewBank(code string, name string) (*Bank, error) {
	bank := Bank{Code: code, Name: name}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()

	err := bank.validator()

	if err != nil {
		return nil, err
	}
	return &bank, nil
}
