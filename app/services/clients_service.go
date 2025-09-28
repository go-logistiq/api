package services

import (
	"strings"
	"sync"

	"github.com/go-logistiq/api/app/models"
	"github.com/go-logistiq/api/app/queries"
	"github.com/go-raptor/raptor/v4"
	"github.com/go-raptor/raptor/v4/errs"
)

type ClientsService struct {
	raptor.Service

	Groups *GroupsService
	DB     *DatabaseService

	lock        sync.RWMutex
	clientCache map[string]int // Cache for client IDs by NATS subject
}

func NewClientsService() *ClientsService {
	return &ClientsService{
		clientCache: make(map[string]int),
	}
}

func (gs *ClientsService) All() (models.Clients, error) {
	return Select[models.Client](gs.DB, queries.GetClients())
}

func (gs *ClientsService) GetBySlug(groupSlug, clientSlug string) (models.Client, error) {
	group, err := gs.Groups.GetBySlug(groupSlug)
	if err != nil {
		return models.Client{}, err
	}

	return SelectOne[models.Client](gs.DB, queries.GetClientBySlug(group.ID, clientSlug))
}

// GetBySubject retrieves a client by NATS subject like "logs.groupSlug.clientSlug"
func (gs *ClientsService) GetIDBySubject(subject string) (int, error) {
	gs.lock.RLock()
	if id, ok := gs.clientCache[subject]; ok {
		gs.lock.RUnlock()
		return id, nil
	}
	gs.lock.RUnlock()

	parts := strings.Split(subject, ".")
	if len(parts) != 3 {
		return 0, errs.NewErrorBadRequest("Invalid subject format")
	}

	groupSlug := parts[1]
	clientSlug := parts[2]

	client, err := gs.GetBySlug(groupSlug, clientSlug)
	if err != nil {
		return 0, err
	}

	gs.lock.Lock()
	gs.clientCache[subject] = client.ID
	gs.lock.Unlock()

	return client.ID, nil
}
