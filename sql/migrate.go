package sql

import (
	"embed"
	"liquid/pkg/db"

	"github.com/pressly/goose/v3"
	log "github.com/sirupsen/logrus"
)

//go:embed *.sql
var embedMigrations embed.FS

func MigrateUp() {
	log.Info("Migrate up")
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}
	if err := goose.Up(db.Instance().DB, "."); err != nil {
		panic(err)
	}
}

func MigrateDown() {
	log.Info("Migrate down")
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}
	if err := goose.DownTo(db.Instance().DB, ".", 0); err != nil {
		panic(err)
	}
}
