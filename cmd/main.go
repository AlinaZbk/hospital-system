package main

import (
	"log"

	"github.com/AlinaZbk/hospital-system/internal/handler"
	"github.com/AlinaZbk/hospital-system/internal/repository"
	"github.com/AlinaZbk/hospital-system/internal/server"
	"github.com/AlinaZbk/hospital-system/internal/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurres while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
