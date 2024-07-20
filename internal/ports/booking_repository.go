package ports

import "booking-service/internal/domain"

type BookingRepository interface {
	Save(booking *domain.Booking) error
	FindByID(id string) (*domain.Booking, error)
	Update(booking *domain.Booking) error
	Delete(id string) error
}
