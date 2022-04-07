package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type PixKeyRepositoryInterface interface {
	ResisterKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	FindAccount(id string) (*Account, error)
}
type PixKey struct {
	Base      `valid:"required"`
	Kind      string   `json:"kind" valid:"notnull"`
	Key       string   `json:"key" valid:"notnull"`
	AccountID string   `json:"accountID" valid:"notnull"`
	Account   *Account `valid:"-"`
	Status    string   `json:"status"`
}

func (pixkey *PixKey) isValid() error {
	_, err := govalidator.ValidateStruct(pixkey)
	if err != nil {
		return err
	}
	return nil
}

func NewPixKey(kind string, key string, account *Account) (*PixKey, error) {
	pixKey := PixKey{Kind: kind, Key: key, Account: account, Status: "active"}

	pixKey.AccountID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()

	err := pixKey.isValid()
	if err != nil {
		return nil, err
	}
	return &pixKey, nil
}
