package systems

import (
	"context"
	"github.com/levpaul/idolscape-backend/internal/core"
	"github.com/levpaul/idolscape-backend/internal/ecs"
)

// PlayerMovementSystem is repsonsible for reading current inputs and moving player
// and updating interest mapping
type PlayerMovementSystem struct {
	BaseSystem
}

func (pm *PlayerMovementSystem) Init()                                        {}
func (pm *PlayerMovementSystem) Update(ctx context.Context, dt core.GameTick) {}
func (pm *PlayerMovementSystem) Add(en ecs.Entity)                            {}
func (pm *PlayerMovementSystem) Remove(en core.EntityID)                      {}
