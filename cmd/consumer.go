package main

import (
	"context"
	"github.com/pasha1coil/order_data_using_nats/configs"
	"github.com/pasha1coil/order_data_using_nats/internal/handler"
	"github.com/pasha1coil/order_data_using_nats/internal/repository"
	database "github.com/pasha1coil/order_data_using_nats/internal/repository/db"
	"github.com/pasha1coil/order_data_using_nats/internal/repository/db/dbmodel"
	"github.com/pasha1coil/order_data_using_nats/internal/server"
	"github.com/pasha1coil/order_data_using_nats/internal/service"
	"github.com/pasha1coil/order_data_using_nats/internal/service/consumer"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/stan.go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	log.SetFormatter(new(log.JSONFormatter))
	if err := configs.InitConfig(); err != nil {
		log.Errorf("error intializing configs: %s", err.Error())
	}

	db, err := database.Mig(dbmodel.DbConfig{
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

	log.Infoln("Init repository")
	rep := repository.InitRepository(db)

	log.Infoln("Init service")
	svc := service.NewService(rep)

	log.Infoln("Init handler")
	handlers := handler.NewHandler(svc)

	log.Infoln("Try restore cache")
	err = svc.RestoreCache()
	if err != nil {
		log.Errorf("Error restore cache: %s", err.Error())
	}

	connect := consumer.CreateSub(*svc)

	log.Infoln("Connecting to STAN")
	err = connect.Connect(
		viper.GetString("nats.cluster_id"),
		viper.GetString("nats.client_id"),
		viper.GetString("nats.host")+":"+viper.GetString("nats.port"))

	if err != nil {
		log.Errorf("Error connecting to STAN: %s", err.Error())
	}
	defer connect.Close()

	log.Infoln("Subscribe to NATS channel")
	sub, err := connect.SubscribeToChannel(viper.GetString("nats.channel"), stan.StartWithLastReceived())

	if err != nil {
		log.Printf("Error subscribing to channel : %s", err.Error())
	}
	//defer sub.Unsubscribe()

	srv := new(server.Server)
	go func() {
		if err := srv.Run(viper.GetString("httpsrv.port"), handlers.InitRoutes()); err != nil {
			log.Errorf("error occured while running http server: %s", err.Error())
		}
	}()

	log.Infoln("App Has Been Activated")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Infoln("App Has Been Downed")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Errorf("an error occurred while shutting down the server:%s", err.Error())
	}
	if err := db.Close(); err != nil {
		log.Errorf("an error occurred while closing the database connection: %s", err.Error())
	}
	if err := sub.Unsubscribe(); err != nil {
		log.Errorf("an error occured while Unsubscribe")
	}
}
