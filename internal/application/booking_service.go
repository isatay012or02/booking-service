package application

import (
	"booking-service/internal/domain"
	"booking-service/internal/ports"
	"time"
)

type BookingServiceImpl struct {
	repo ports.BookingRepository
}

func NewBookingService(repo ports.BookingRepository) ports.BookingService {
	return &BookingServiceImpl{repo: repo}
}

func (s *BookingServiceImpl) CreateBooking(userID, flightID string) (*domain.Booking, error) {
	booking := &domain.Booking{
		ID:        generateID(),
		UserID:    userID,
		FlightID:  flightID,
		Status:    domain.Pending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := s.repo.Save(booking)
	if err != nil {
		return nil, err
	}

	return booking, nil
}

func (s *BookingServiceImpl) GetBooking(id string) (*domain.Booking, error) {

	resp, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *BookingServiceImpl) UpdateBookingStatus(id string, status domain.BookingStatus) (*domain.Booking, error) {
	booking, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	booking.Status = status
	booking.UpdatedAt = time.Now()
	err = s.repo.Update(booking)
	if err != nil {
		return nil, err
	}

	return booking, nil
}

func (s *BookingServiceImpl) CancelBooking(id string) error {
	booking, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	booking.Status = domain.Cancelled
	booking.UpdatedAt = time.Now()

	return s.repo.Update(booking)
}

func generateID() string {
	return "some-unique-id" // Реализуйте генерацию уникального ID
}
