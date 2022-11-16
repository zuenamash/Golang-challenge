package processors

import (
	"sync"
	"fmt"
) 

type DriverRanking struct {
	// Fill in your properties here
	driver *Driver
	rank float64

}

func (*DriverRanking) String() string {
	// Implement this function
	return fmt.Sprintf("%v Rating-%2f",
d.driver.Name,d.rank)
}

type HotelRanking struct {
	// Fill in your properties here
	hotel *Hotel
	rank float64
}

func (*HotelRanking) String() string {
	// Implement this function
	return ""
}

type ProcessorInterface interface {
	StartProcessing() error
	GetTopRankedDriver() *DriverRanking
	GetTopRankedHotel() *HotelRanking
}

func CreateProcessorFromData(data *TripsData, wg *sync.WaitGroup) ProcessorInterface  {
	// @todo Initialize your processor here
	return nil
}
