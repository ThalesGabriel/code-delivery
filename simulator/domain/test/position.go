package model

import (
	"testing"

	"github.com/ThalesGabriel/code-delivery/simulator/domain/model"
	"github.com/stretchr/testify/require"
)

func TestModel_Position(t *testing.T) {
	lat := -15.82594
	long := -47.92923

	_, err := model.NewPosition(lat, long)
	require.NotNil(t, err)
}
