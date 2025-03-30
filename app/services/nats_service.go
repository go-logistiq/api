package services

import (
	"time"

	"github.com/go-raptor/raptor/v3"
	"github.com/nats-io/nats.go"
)

type NATSService struct {
	raptor.Service

	natsURL  string
	natsConn *nats.Conn

	Worker *WorkerService
}

func NewNATSService(c *raptor.Config) *NATSService {
	natsURL := c.AppConfig["nats_url"]
	if natsURL == "" {
		natsURL = nats.DefaultURL
	}

	ns := &NATSService{
		natsURL: natsURL,
	}

	ns.OnInit(ns.Init)
	ns.OnShutdown(ns.Shutdown)

	return ns
}

func (ns *NATSService) Init() error {
	var err error
	ns.natsConn, err = nats.Connect(ns.natsURL,
		nats.MaxReconnects(-1),
		nats.ReconnectWait(5*time.Second),
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			ns.Log.Warn("NATS disconnected", "error", err)
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			ns.Log.Warn("NATS reconnected", "url", nc.ConnectedUrl())
		}),
	)
	if err != nil {
		return err
	}

	return ns.subscribeToLogs()
}

func (ls *NATSService) Shutdown() error {
	if ls.natsConn != nil {
		err := ls.natsConn.FlushTimeout(5 * time.Second)
		ls.natsConn.Close()
		return err
	}
	return nil
}

func (ls *NATSService) subscribeToLogs() error {
	_, err := ls.natsConn.Subscribe("logs.*.*", func(msg *nats.Msg) {
		ls.Worker.MessageChan <- msg
	})

	return err
}
