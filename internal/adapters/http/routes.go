package http

func (s *server) routers() {

	v1 := s.router.Group("/api/v1")
	{
		v1.POST("booking", s.handlers.CreateBooking)
		v1.GET("booking", s.handlers.GetBooking)
		v1.PUT("booking", s.handlers.UpdateBookingStatus)
	}

	s.router.GET("/health")
}
