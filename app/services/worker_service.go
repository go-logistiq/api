package services

import (
	"strconv"
	"sync"

	"github.com/go-raptor/raptor/v3"
	"github.com/nats-io/nats.go"
)

type WorkerService struct {
	raptor.Service

	Logs *LogsService

	numWorkers  int
	MessageChan chan *nats.Msg
	wg          sync.WaitGroup
}

func NewWorkerService(c *raptor.Config) *WorkerService {
	numWorkers, err := strconv.Atoi(c.AppConfig["workers"])
	if err != nil {
		numWorkers = 4
	}

	ws := &WorkerService{
		numWorkers:  numWorkers,
		MessageChan: make(chan *nats.Msg, 1000),
	}

	ws.OnInit(ws.Init)
	ws.OnShutdown(ws.Shutdown)

	return ws
}

func (ws *WorkerService) Init() error {
	for i := 0; i < ws.numWorkers; i++ {
		go ws.startWorker()
	}

	ws.Log.Info("Worker service started", "num_workers", ws.numWorkers)
	return nil
}

func (ws *WorkerService) Shutdown() error {
	ws.Log.Info("Worker service shutting down")
	close(ws.MessageChan)
	ws.wg.Wait()
	ws.Log.Info("Worker service shut down")
	return nil
}

func (ws *WorkerService) startWorker() {
	ws.wg.Add(1)
	defer ws.wg.Done()

	for msg := range ws.MessageChan {
		logs, err := ws.Logs.ParseNATSMessage(msg)
		if err != nil {
			ws.Log.Error("failed to parse NATS message", "error", err)
			continue
		}
		ws.Logs.Save(logs)
	}
}
