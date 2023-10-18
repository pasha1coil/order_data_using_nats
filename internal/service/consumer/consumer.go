package consumer

import (
	"github.com/nats-io/stan.go"
	"github.com/pasha1coil/order_data_using_nats/internal/service"
	log "github.com/sirupsen/logrus"
)

type StanSubscriber struct {
	con stan.Conn
	svc service.Service
}

func CreateSub(svc service.Service) *StanSubscriber {
	log.Infoln("Create sub, message from consumer")
	subscriber := StanSubscriber{
		svc: svc,
	}
	return &subscriber
}

func (s *StanSubscriber) Connect(clusterID string, clientID string, URL string) error {
	log.Infoln("Connect, message from consumer")
	con, err := stan.Connect(clusterID, clientID, stan.NatsURL(URL))
	if err != nil {
		return err
	}
	s.con = con
	return nil
}

func (s *StanSubscriber) Close() {
	log.Infoln("Close, message from consumer")
	if s.con != nil {
		s.con.Close()
	}
}

func (s *StanSubscriber) SubscribeToChannel(channel string, opts ...stan.SubscriptionOption) (stan.Subscription, error) {
	log.Infoln("Subscribe to channel, message from consumer")
	sub, err := s.con.Subscribe(channel, s.handlerMsg, opts...)
	if err != nil {
		log.Errorf("Can't connect, message from consumer: %s", err.Error())
		s.svc.RestoreCache()
	}
	return sub, err
}

func (s *StanSubscriber) handlerMsg(msg *stan.Msg) {
	log.Infoln("Received a new message from NATS, message from consumer")
	err := s.svc.SaveOrderData(msg.Data)
	if err != nil {
		log.Errorf("error saving: %s", err.Error())
	}
}
