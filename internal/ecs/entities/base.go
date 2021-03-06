package entities

import (
	"github.com/levpaul/gecserv/internal/core"
	"sync"
)

var (
	idCounter core.EntityID
	counterMu sync.Mutex
)

type BaseEntity struct {
	id core.EntityID
}

func (e *BaseEntity) ID() core.EntityID {
	return e.id
}

func NewBaseEntity() *BaseEntity {
	be := new(BaseEntity)
	be.id = newEntityID()
	return be
}

func newEntityID() core.EntityID {
	counterMu.Lock()
	defer counterMu.Unlock()
	idCounter += 1
	return idCounter
}
