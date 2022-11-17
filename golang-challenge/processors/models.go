package processors

// Hotel ...
type Hotel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Driver ...
type Driver struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Trip ...
type Trip struct {
	ID           string  `json:"id"`
	DriverID     string  `json:"driver_id"`
	HotelID      string  `json:"hotel_id"`
	HotelRating  float64 `json:"hotel_rating"`
	DriverRating float64 `json:"driver_rating"`
	Status       string  `json:"status"`
	Driver       *Driver `json:"-"`
	Hotel        *Hotel  `json:"-"`
}

// TripsData ...
type TripsData struct {
	Drivers []*Driver  `json:"drivers"`
	Hotels  []*Hotel   `json:"hotels"`
	Trips   chan *Trip `json:"-"`
}
