package producer

import (
	"fmt"
	"math"
	"math/rand/v2"
	"strconv"
	"time"
)

func Logf(msg string, a ...interface{}) {
	fmt.Printf(msg, a...)
	fmt.Println()
}


type TruckData struct {
	TruckID string
	Latitude float64
	Longitude float64
	Speed float64
	FuelLevel float64
	Temperature float64
	LastUpdated time.Time
}

type TruckMessage struct {
	TruckID string `json:"truck_id"`
	Timestamp time.Time `json:"timestamp"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Speed float64 `json:"speed"`
	FuelLevel float64 `json:"fuel_level"`
	Temperature float64 `json:"temperature"`
}

func NewTruckState(id int) *TruckData {
	return &TruckData{
		TruckID: strconv.Itoa(id),
		Latitude: 48.8566 + rand.Float64()/10,
		Longitude: 2.3522 + rand.Float64()/10,
		Speed: 50 + rand.Float64()*10,
		FuelLevel: 100.0,
		Temperature: 65 + rand.Float64()*5,
		LastUpdated: time.Now(),
	}
}

func (t *TruckData) Update() {
	t.Latitude += (rand.Float64() - 0.5) * 0.01
	t.Longitude += (rand.Float64() - 0.5) * 0.01

	deltaSpeed := (rand.Float64() - 0.5) * 5
	t.Speed = math.Max(0, math.Min(120, t.Speed+deltaSpeed))

	t.FuelLevel = math.Max(0, t.FuelLevel - t.Speed*0.005)

	t.Temperature += (rand.Float64() - 0.5) * 2
	t.LastUpdated = time.Now()
}

func (t *TruckData) ToMessage() TruckMessage {
	return TruckMessage{
		TruckID: t.TruckID,
		Timestamp: t.LastUpdated,
		Latitude: t.Latitude,
		Longitude: t.Longitude,
		Speed: t.Speed,
		FuelLevel: t.FuelLevel,
		Temperature: t.Temperature,
	}
}

func MakeTrucks(num_trucks int) []*TruckData {
	truck_data := make([]*TruckData, num_trucks)
	for i := 0; i < num_trucks; i++ {
		truck_data[i] = NewTruckState(i)
	}
	return truck_data
}



