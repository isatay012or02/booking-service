package ports

import "booking-service/internal/domain"

type BookingService interface {
	CreateBooking(userID, flightID string) (*domain.Booking, error)
	GetBooking(id string) (*domain.Booking, error)
	UpdateBookingStatus(id string, status domain.BookingStatus) (*domain.Booking, error)
	CancelBooking(id string) error
}
