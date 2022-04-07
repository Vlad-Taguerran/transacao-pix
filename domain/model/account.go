package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Account struct {
	Base      `valid:"required"`
	OwnerName string    `json:"owner_name" valid:"notnull"`
	Bank      *Bank     `valid:"-"`
	Number    string    `json:"number" valid:"required"`
	PixKeys   []*PixKey `valid:"-"`
}

func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)
	if err != nil {
		return err
	}
	return nil
}

func NewAccount(bank *Bank, ownerName string, number string) (*Account, error) {
	accont := Account{Bank: nil, OwnerName: ownerName, Number: number}

	accont.ID = uuid.NewV4().String()
	accont.CreatedAt = time.Now()

	err := accont.isValid()
	if err != nil {
		return nil, err
	}
	return &accont, nil
}
