package repository

import (
	"booking-service/internal/domain"
	"booking-service/internal/ports"
	"gorm.io/gorm"
)

type BookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) ports.BookingRepository {
	return &BookingRepository{db: db}
}

func (r *BookingRepository) Save(booking *domain.Booking) error {
	return r.db.Create(booking).Error
}

func (r *BookingRepository) FindByID(id string) (*domain.Booking, error) {
	var booking domain.Booking
	result := r.db.First(&booking, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &booking, nil
}

func (r *BookingRepository) Update(booking *domain.Booking) error {
	return r.db.Save(booking).Error
}

func (r *BookingRepository) Delete(id string) error {
	return r.db.Delete(&domain.Booking{}, "id = ?", id).Error
}
