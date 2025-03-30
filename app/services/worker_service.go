package services

import (
	"sync"

	"github.com/go-raptor/raptor/v3"
	"github.com/nats-io/nats.go"
)

type WorkerService struct {
	raptor.Service

	numWorkers  int
	MessageChan chan *nats.Msg
	wg          sync.WaitGroup
}

func NewWorkerService(numWorkers int) *WorkerService {
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
	close(ws.MessageChan)
	ws.wg.Wait()
	return nil
}

func (ws *WorkerService) startWorker() {
	ws.wg.Add(1)
	defer ws.wg.Done()

	for msg := range ws.MessageChan {
		ws.Log.Info("Processing message", "subject", msg.Subject, "data", string(msg.Data))
	}
}
