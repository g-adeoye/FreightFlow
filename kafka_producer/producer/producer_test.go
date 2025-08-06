package producer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestNewTruckState(t *testing.T) {
	truck_state := NewTruckState(3)
	t.Run("asserting Fuel level is 100", func(t *testing.T) {
		assert.Equal(t, truck_state.FuelLevel, 100.0)
	})
	t.Run("asserting id is 3", func(t *testing.T) {
		assert.Equal(t, truck_state.TruckID, "3")
	})
}

func TestMakeTrucks(t *testing.T) {
	truck_data := MakeTrucks(5)
	t.Run(
		"asserting number of trucks is 5",
		func(t *testing.T) {
			assert.Equal(t, 5, len(truck_data))
		},
	)
}

