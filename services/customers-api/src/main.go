package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Astak/otus-docker-basics-homework/web-service-gin/config"
	"github.com/Astak/otus-docker-basics-homework/web-service-gin/handler"
	"github.com/Astak/otus-docker-basics-homework/web-service-gin/migrations"
	"github.com/Astak/otus-docker-basics-homework/web-service-gin/router"
	"github.com/rs/zerolog/log"
)

func main() {
	configPath := flag.String("configpath", "", "Config Path")
	migration := flag.Bool("migration", false, "Migration")
	flag.Parse()
	if configPath == nil || len(*configPath) == 0 {
		log.Fatal().Msgf("Unable to load config path. Empty path specified. %s", *configPath)
	}
	if _, err := os.Stat(*configPath); os.IsNotExist(err) {
		log.Fatal().Msgf("Unable to load config path. Path not found. %s", *configPath)
	}
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal().Msgf(err.Error())
	}
	if *migration {
		migrations.Migrate(cfg)
	} else {
		h := handler.LoadHandlerFromConfig(cfg)
		r := router.SetupRouter(h, cfg)
		r.Run(fmt.Sprintf("0.0.0.0:%d", cfg.Port))
	}
}
