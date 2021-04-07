package model

import (
	"time"

	validator "github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Route struct {
	ID        string     `json:"routeId"`
	ClientId  string     `json:"clientId"`
	Positions []Position `json:"position"`
	CreatedAt time.Time  `json:"createdAt"`
}

func (route *Route) isValid() error {
	_, err := validator.ValidateStruct(route)
	if err != nil {
		return err
	}
	return nil
}

func NewRoute(clientId string) (Route, error) {
	route := Route{
		ID:        uuid.NewV4().String(),
		ClientId:  clientId,
		CreatedAt: time.Now(),
	}

	err := route.isValid()
	if err != nil {
		return Route{}, err
	}

	return route, nil
}
