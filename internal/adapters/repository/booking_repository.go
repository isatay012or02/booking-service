package repository

import (
	"booking-service/internal/domain"
	"booking-service/internal/ports"
	"gorm.io/gorm"
)

type PostgresBookingRepository struct {
	db *gorm.DB
}

func NewPostgresBookingRepository(db *gorm.DB) ports.BookingRepository {
	return &PostgresBookingRepository{db: db}
}

func (r *PostgresBookingRepository) Save(booking *domain.Booking) error {
	return r.db.Create(booking).Error
}

func (r *PostgresBookingRepository) FindByID(id string) (*domain.Booking, error) {
	var booking domain.Booking
	result := r.db.First(&booking, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &booking, nil
}

func (r *PostgresBookingRepository) Update(booking *domain.Booking) error {
	return r.db.Save(booking).Error
}

func (r *PostgresBookingRepository) Delete(id string) error {
	return r.db.Delete(&domain.Booking{}, "id = ?", id).Error
}
