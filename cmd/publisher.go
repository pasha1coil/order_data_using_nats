package main

import (
	"github.com/pasha1coil/order_data_using_nats/configs"
	"github.com/pasha1coil/order_data_using_nats/internal/service/publisher"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	log.SetFormatter(new(log.JSONFormatter))
	if err := configs.InitConfig(); err != nil {
		log.Errorf("error intializing configs: %s", err.Error())
	}
	client := publisher.CreateSTAN()
	err := client.Connect(viper.GetString("nats.cluster_id"),
		viper.GetString("nats.publisher_id"),
		viper.GetString("nats.host")+":"+viper.GetString("nats.port"))
	defer client.Close()
	if err != nil {
		log.Fatalf("Error connecting to nats : %s", err.Error())
	}
	_ = client.PublishFromCLI(viper.GetString("nats.channel"))
}
