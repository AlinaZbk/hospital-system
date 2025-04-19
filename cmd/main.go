package main

import (
	"log"

	"github.com/AlinaZbk/hospital-system/internal/handler"
	"github.com/AlinaZbk/hospital-system/internal/server"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(server.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurres while running http server: %s", err.Error())
	}
}
