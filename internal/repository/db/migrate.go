package database

import (
	"fmt"
	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pasha1coil/order_data_using_nats/internal/repository/db/dbmodel"
	log "github.com/sirupsen/logrus"
)

func Mig(cfg dbmodel.DbConfig) (*sqlx.DB, error) {
	log.Infoln("Migrate database")
	db, err := InitDB(cfg)
	if err != nil {
		return nil, err
	}

	mig, err := migrate.New(
		"file://internal/repository/db/sql",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			cfg.Uname, cfg.Pass, cfg.Host, cfg.Port, cfg.NameDB, cfg.SSL))
	if err != nil {
		log.Errorf("Error New migrate:%s", err.Error())
		return nil, err
	}

	err = mig.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Errorf("Error Up migrate:%s", err.Error())
		return nil, err
	}
	return db, nil

}
