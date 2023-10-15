package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"wbl0/internal/repository"
	database "wbl0/internal/repository/db"
)

func main() {

	//worker := worker.InitApp(*config)
	//worker.Run()
	log.SetFormatter(new(log.JSONFormatter))
	if err := initConfig(); err != nil {
		log.Errorf("error intializing configs: %s", err.Error())
	}
	db, err := database.InitDB(database.DbConfig{
		Host:       viper.GetString("db.Host"),
		Port:       viper.GetString("db.Port"),
		SSL:        viper.GetString("db.SSL"),
		Uname:      viper.GetString("db.Uname"),
		Pass:       viper.GetString("db.Pass"),
		NameDB:     viper.GetString("db.NameDB"),
		DriverName: viper.GetString("db.DriverName"),
	})
	if err != nil {
		log.Errorf("failed to initialize db: %s", err.Error())
	}
	rep := repository.InitRepository(db)

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
