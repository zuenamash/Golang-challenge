package processors

import (
	"log"
	"sync"
	"sync/atomic"
)

// Processor should implement ProcessorInterface
// Hint: Try using goroutines to process the data in parallel
// Requirement: Do not store all Trips in memory and cache your results
type Processor struct {
	data      *TripsData
	wg        *sync.WaitGroup
	topDriver *DriverRanking
	topHotel  *HotelRanking

	mu      sync.Mutex
	drivers map[string]*DriverRanking
	hotels  map[string]*HotelRanking
}

func (p *Processor) addTrip(trip *Trip) {
	p.mu.Lock()
	defer p.mu.Unlock()
	driverRank, ok := p.drivers[trip.DriverID]
	if !ok {
		driverRank = &DriverRanking{
			Driver:     trip.Driver,
			TotalTrips: 1,
			Rating:     trip.DriverRating,
		}
	} else {
		driverRank.TotalTrips++
		driverRank.Rating = (driverRank.Rating + trip.DriverRating) / 2
	}

	p.drivers[trip.DriverID] = driverRank

	hotelRank, ok := p.hotels[trip.HotelID]
	if !ok {
		hotelRank = &HotelRanking{
			Hotel:      trip.Hotel,
			TotalTrips: 1,
			Rating:     trip.HotelRating,
		}
	} else {
		hotelRank.TotalTrips++
		hotelRank.Rating = (hotelRank.Rating + trip.HotelRating) / 2
	}

	p.hotels[trip.HotelID] = hotelRank
}

// StartProcessing should start processing the data
func (p *Processor) StartProcessing() error {
	p.drivers = make(map[string]*DriverRanking)
	p.hotels = make(map[string]*HotelRanking)

	var tripCount uint64

	// create 10 workers to process the data
	for i := 0; i < 10; i++ {
		p.wg.Add(1)
		log.Println("starting worker", i)

		go func(worker int) {
			for trip := range p.data.Trips {
				atomic.AddUint64(&tripCount, 1)
				log.Println("worker", worker, "processing trip", tripCount)
				p.addTrip(trip)
			}
			p.wg.Done()
		}(i)
	}

	var topDriver *DriverRanking
	for _, driver := range p.drivers {
		if topDriver == nil || driver.Rating > topDriver.Rating {
			topDriver = driver
		}
	}
	p.topDriver = topDriver

	var topHotel *HotelRanking
	for _, hotel := range p.hotels {
		if topHotel == nil || hotel.Rating > topHotel.Rating {
			topHotel = hotel
		}
	}
	p.topHotel = topHotel

	p.wg.Done()
	return nil
}

// GetTopRankedDriver should return the top ranked driver
func (p *Processor) GetTopRankedDriver() *DriverRanking {
	return p.topDriver
}

// GetTopRankedHotel should return the top ranked hotel
func (p *Processor) GetTopRankedHotel() *HotelRanking {
	return p.topHotel
}
