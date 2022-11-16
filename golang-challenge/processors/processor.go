go package processors

import(
	"sync"
)
// Processor should implement ProcessorInterface
// Hint: Try using goroutines to process the data in parallel
// Requirement: Do not store all Trips in memory and cache your results

type Processor struct {
	data *TripsData
	wg *sync.WaitGroup

}

func (p Processor)StartProcessing()error{
	return nil
}
func(p Processor)GetTopRankedHotel()
*HotelRanking {
	return nil
}