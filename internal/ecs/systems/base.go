package systems

import (
	"github.com/levpaul/gecserv/internal/ecs"
)

type BaseSystem struct {
	sa *ecs.SectorAdmin
}

func (bs *BaseSystem) SetSectorAdmin(sa *ecs.SectorAdmin) {
	bs.sa = sa
}

func (bs *BaseSystem) GetSectorAdmin() *ecs.SectorAdmin {
	return bs.sa
}
