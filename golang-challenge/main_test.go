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

results := DriverAllrates(trips)
fmt.Println("mapping------------", results)
drivers := Driverrates(results)
fmt.Println("drivers and their rates ------------", driver)
bestdriver := findBestDriver(drivers)

// best driver :=findBestDriver(trips)
fmt.Printf("The best driver is %s with a rate of %.if with %d number of trips \n", bestdriver.Name, bestdriver.rate
bestdriver.Counter)
// averagerate := findAvaragerating(trips)
// fmt.printf("Average rate is %d, averagerate")

}
