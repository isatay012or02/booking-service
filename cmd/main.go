package main

import (
	"booking-service/config"
	"booking-service/internal/adapters/http"
	"booking-service/internal/adapters/repository"
	"booking-service/internal/application"
	"booking-service/internal/handlers"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	configuration, err := config.Init("config.json")
	if err != nil {
		panic(err)
	}

	repoDB, err := repository.NewDB(configuration.DB)
	if err != nil {
		panic(err)
	}

	repo, err := repository.NewBookingRepository(repoDB)
	if err != nil {
		panic(err)
	}

	service := application.NewBookingService(repo)

	handler := handlers.NewBookingHandler(service)

	srv, err := http.NewServer(configuration, handler)
	if err != nil {
		panic(err)
	}

	startServerErrorCH := srv.Start()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err = <-startServerErrorCH:
		{
			panic(err)
		}
	case q := <-quit:
		{
			fmt.Printf("receive signal %s, stopping server...\n", q.String())
			//appLoger.ServerInfo("main", fmt.Sprintf("receive signal %s, stopping server...\n", q.String()))
			if err = srv.Stop(); err != nil {
				fmt.Printf("stop server error: %s\n", err.Error())
				//appLoger.ServerError("main", err.Error(), "stop server error")
			}
		}
	}
}
