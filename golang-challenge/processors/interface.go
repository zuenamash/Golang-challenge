package processors

import (
	"fmt"
	"sync"
)

// DriverRanking ...
type DriverRanking struct {
	Driver     *Driver
	TotalTrips int
	Rating     float64
}

func (d *DriverRanking) String() string {
	return fmt.Sprintf("Driver: %s, Rating: %f", d.Driver.Name, d.Rating)
}

// HotelRanking ...
type HotelRanking struct {
	Hotel      *Hotel
	TotalTrips int
	Rating     float64
}

func (h *HotelRanking) String() string {
	return fmt.Sprintf("Hotel: %s, Rating: %f", h.Hotel.Name, h.Rating)
}

// ProcessorInterface ...
type ProcessorInterface interface {
	StartProcessing() error
	GetTopRankedDriver() *DriverRanking
	GetTopRankedHotel() *HotelRanking
}

// CreateProcessorFromData ...
func CreateProcessorFromData(data *TripsData, wg *sync.WaitGroup) ProcessorInterface {
	p := Processor{
		data: data,
		wg:   wg,
	}
	return &p
}
