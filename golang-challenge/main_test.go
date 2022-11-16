// Implement the following tests:
// Ensure top driver is found the processor
// Ensure top hotel is found the processor
// Ensure memory consumption is less than 128 mb

package main
import "fmt"

type Trip struct {
	Driver string
	Hotel string
	Rate int
}
type Driver struct {
	Name string
	Rate int
}
func main() {
	trips := []Trip{}
		bestdriver := findBestDriver(trips)
		fmt.Print("The best driver is %s with a rate of %d",bestdriver.Name, bestdriver.Rate)
		averagerate := findAveragerating(trips)
		fmt.Printf("Average rate is %d", averagerate)
}
func findBestDriver(data []Trip) *Driver {
		maxrate := 0
		bestdriver := &Driver{}
		for_, driver := rage data {
			if driver.Rate > maxrate {
				bestdriver.Name = driver.Driver
				bestdriver.Rate = driver.Rate
			}
		}
		return bestdriver
	}
func findAveragerating(data []Trip) int {
		rates := 0
		for _, driver := rage data {
			rates += driver.Rate
		}
		return rates / len{data}
	}	
