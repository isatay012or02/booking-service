package handlers

import (
	"booking-service/internal/common"
	"booking-service/internal/domain"
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

// CreateBooking godoc
// @Summary Создать букинг
// @Description Создания букинга
// @Tags BookingV1
// @Accept json
// @Produce json
// @Param userID query string true "ID пользователя"
// @Param flightID query string true "ID рейса"
// @Success 200 {object} common.BaseResponse{}
// @Failure 400,422 {object} object{resultCode=int,resultDescription=string,data=object}
// @Router /api/v1/booking [post]
func (h *BookingHandler) CreateBooking(ctx *gin.Context) {
	var req request.CreateBooking
	if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, common.BaseResponse{Data: err.Error()})
		return
	}

	booking, err := h.service.CreateBooking(req.UserID, req.FlightID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.BaseResponse{Data: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, common.BaseResponse{Data: booking})
	return
}

// GetBooking godoc
// @Summary Получить букинг
// @Description Получения букинга
// @Tags BookingV1
// @Accept json
// @Produce json
// @Param ID query string true "ID букинга"
// @Success 200 {object} common.BaseResponse{}
// @Failure 400,422 {object} object{resultCode=int,resultDescription=string,data=object}
// @Router /api/v1/booking [get]
func (h *BookingHandler) GetBooking(ctx *gin.Context) {
	id := ctx.Query("id")

	booking, err := h.service.GetBooking(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.BaseResponse{Data: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, common.BaseResponse{Data: booking})
	return
}

// UpdateBookingStatus godoc
// @Summary Обновить букинг
// @Description Обновдления букинга
// @Tags BookingV1
// @Accept json
// @Produce json
// @Param ID query string true "ID букинга"
// @Param status query string true "Статус букинга"
// @Success 200 {object} common.BaseResponse{}
// @Failure 400,422 {object} object{resultCode=int,resultDescription=string,data=object}
// @Router /api/v1/booking [put]
func (h *BookingHandler) UpdateBookingStatus(ctx *gin.Context) {
	id := ctx.Query("id")
	status := domain.BookingStatus(ctx.Query("status"))

	booking, err := h.service.UpdateBookingStatus(id, status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.BaseResponse{Data: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, common.BaseResponse{Data: booking})
	return
}

func (h *BookingHandler) CancelBooking(ctx *gin.Context) {
	id := ctx.Query("id")

	err := h.service.CancelBooking(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.BaseResponse{Data: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, common.BaseResponse{})
	return
}
