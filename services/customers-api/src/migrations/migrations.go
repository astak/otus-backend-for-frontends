package migrations

import (
	"log"

	"github.com/Astak/otus-docker-basics-homework/web-service-gin/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(config config.Config) {
	dbUrl := config.GetDbDevUrl()
	m, err := migrate.New("file://migrations/sql", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
