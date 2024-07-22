package http

func (s *server) routers() {

	v1 := s.router.Group("/api/v1")
	{
		v1.POST("booking", s.handlers.CreateBooking)
	}

	s.router.GET("/health")
}
