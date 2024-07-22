package request

type CreateBooking struct {
	UserID   string `json:"user_id"`
	FlightID string `json:"flight_id"`
}
