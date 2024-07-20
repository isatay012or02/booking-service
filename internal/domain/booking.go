package domain

import "time"

type BookingStatus string

const (
	Pending   BookingStatus = "Pending"
	Completed BookingStatus = "Completed"
	Cancelled BookingStatus = "Cancelled"
)

type Booking struct {
	ID        string `gorm:"primaryKey"`
	UserID    string
	FlightID  string
	Status    BookingStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}
