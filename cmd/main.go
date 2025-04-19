package main

import (
	"log"

	"github.com/AlinaZbk/hospital-system/internal/handler"
	"github.com/AlinaZbk/hospital-system/internal/repository"
	"github.com/AlinaZbk/hospital-system/internal/server"
	"github.com/AlinaZbk/hospital-system/internal/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	
	srv := new(server.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurres while running http server: %s", err.Error())
	}
}
