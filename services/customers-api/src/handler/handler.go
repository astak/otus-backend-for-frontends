package handler

import (
	"github.com/Astak/otus-docker-basics-homework/web-service-gin/config"
	"github.com/Astak/otus-docker-basics-homework/web-service-gin/data"
)

type Handler struct {
	UserRepo data.ProfileRepository
}

func NewHandler(ur data.ProfileRepository) *Handler {
	return &Handler{UserRepo: ur}
}

func LoadHandler(configPath *string) *Handler {
	cfg, _ := config.LoadConfig(configPath)
	return LoadHandlerFromConfig(cfg)
}

func LoadHandlerFromConfig(cfg config.Config) *Handler {
	database := data.NewDatabase(cfg)
	userRepo, _ := data.NewProfileRepository(database)
	handler := NewHandler(userRepo)
	return handler
}
