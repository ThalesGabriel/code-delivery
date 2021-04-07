package model

import (
	"testing"

	"github.com/ThalesGabriel/code-delivery/simulator/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

// type Route struct {
// 	id        string
// 	clientId  string
// 	positions []Position
// 	createdAt time.Time
// }

func TestModel_Route(t *testing.T) {
	clientId := "1"

	route, err := model.NewRoute(clientId)
	require.NotNil(t, err)

	require.NotEmpty(t, uuid.FromStringOrNil(route.ID))
}
