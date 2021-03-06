package model

import (
	"time"

	validator "github.com/asaskevich/govalidator"
)

type PartialRoutePosition struct {
	ID        string    `json:"routeId"`
	ClientId  string    `json:"clientId"`
	Positions []float64 `json:"positions"`
	Finished  bool      `json:"finished"`
	CreatedAt time.Time `json:"createdAt"`
}

func (partialRoutePosition *PartialRoutePosition) isValid() error {
	_, err := validator.ValidateStruct(partialRoutePosition)
	if err != nil {
		return err
	}
	return nil
}

func NewPartialRoutePosition(ID, clientId string, positions []float64, finished bool) (PartialRoutePosition, error) {
	partialRoutePosition := PartialRoutePosition{
		ID:        ID,
		ClientId:  clientId,
		Positions: positions,
		Finished:  finished,
		CreatedAt: time.Now(),
	}

	err := partialRoutePosition.isValid()
	if err != nil {
		return PartialRoutePosition{}, err
	}

	return partialRoutePosition, nil
}
