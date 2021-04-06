package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/ThalesGabriel/code-delivery/simulator/domain/model"
)

type Route model.Route
type Position model.Position
type PartialRoutePosition model.PartialRoutePosition

// NewRoute creates a *Route struct
func NewRoute() *Route {
	return &Route{}
}

func getCoordinate(data []string, which_coordinate int) (float64, error) {
	coordinate, err := strconv.ParseFloat(data[which_coordinate], 64)
	if err != nil {
		return 0, err
	}
	return coordinate, nil
}

func (r *Route) LoadPositions() error {
	if r.ClientId == "" {
		return errors.New("Route ID not informed")
	}

	destination := "destinations/" + r.ClientId + ".txt"
	f, err := os.Open(destination)
	if err != nil {
		// fmt.Errorf("Could not open file: ", destination)
		return err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		lat, err := getCoordinate(data, 0)
		if err != nil {
			// fmt.Errorf("Could not open file: ", destination)
			return err
		}
		long, err := getCoordinate(data, 1)
		if err != nil {
			// fmt.Errorf("Could not open file: ", destination)
			return err
		}

		new_position, err := model.NewPosition(lat, long)
		if err != nil {
			// fmt.Errorf("Could not open file: ", destination)
			return err
		}

		r.Positions = append(r.Positions, new_position)
	}

	return nil
}

func (r *Route) ExportJsonPositions() ([]string, error) {
	var result []string
	total := len(r.Positions)

	for key, value := range r.Positions {
		clientId := r.ClientId
		position := []float64{value.Lat, value.Long}
		finished := false
		if total-1 == key {
			finished = true
		}
		route, err := model.NewPartialRoutePosition(clientId, position, finished)
		if err != nil {
			return nil, err
		}

		jsonRoute, err := json.Marshal(route)
		if err != nil {
			return nil, err
		}

		result = append(result, string(jsonRoute))
	}

	return result, nil
}
