package handlers

import (
	"booking-service/internal/domain/request"
	"booking-service/internal/ports"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BookingHandler struct {
	service ports.BookingService
}

func NewBookingHandler(service ports.BookingService) *BookingHandler {
	return &BookingHandler{service: service}
}

func (h *BookingHandler) CreateBooking(ctx *gin.Context) {
	var req request.CreateBooking
	if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	booking, err := h.service.CreateBooking(req.UserID, req.FlightID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, booking)
	return
}
