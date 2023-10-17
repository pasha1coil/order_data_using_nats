package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"wbl0/internal/repository/db/dbmodel"
)

func InitDB(cfg dbmodel.DbConfig) (*sqlx.DB, error) {
	log.Infoln("Init database")
	fmt.Println(cfg.DriverName, cfg.Host, cfg.Port, cfg.Uname, cfg.NameDB, cfg.Pass, cfg.SSL)
	db, err := sqlx.Open(cfg.DriverName, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Uname, cfg.NameDB, cfg.Pass, cfg.SSL))
	if err != nil {
		return nil, err
	}
	log.Infoln("Ping database")
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
