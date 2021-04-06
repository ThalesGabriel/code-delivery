package model

import (
	"time"

	validator "github.com/asaskevich/govalidator"
)

type PositionRepositoryInterface interface {
	// NewPosition(pixKey *PixKey) (*PixKey, error)
	// FindKeyByKind(key string, kind string) (*PixKey, error)
	// AddBank(bank *Bank) error
	// AddAccount(account *Account) error
	// FindAccount(id string) (*Account, error)
}

type Position struct {
	Lat       float64   `json:"latitude"`
	Long      float64   `json:"longitude"`
	CreatedAt time.Time `json:"createdAt"`
}

func (position *Position) IsValid() error {
	_, err := validator.ValidateStruct(position)
	if err != nil {
		return err
	}
	return nil
}

func NewPosition(lat, long float64) (Position, error) {
	position := Position{
		Lat:       lat,
		Long:      long,
		CreatedAt: time.Now(),
	}

	err := position.IsValid()
	if err != nil {
		return Position{}, err
	}

	return position, nil
}
