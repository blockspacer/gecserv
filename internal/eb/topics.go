package eb

import (
	"github.com/levpaul/gecserv/internal/core"
	"github.com/levpaul/gecserv/internal/fb"
)

const (
	// Simulation messages
	S_LOGIN = 0 + iota
	S_LOGOUT
	S_GAMETICK_DONE
	// Network messages
	N_CONNECT = 128 + iota
	N_DISCONN
	N_PLAYER_SYNC
	// Used as marker for event topics - insert new topics ABOVE this one
	NUMTOPICS = 255
)

// Bindings of message topics to types
type (
	S_LOGIN_T         *fb.PlayerT
	S_LOGOUT_T        float64
	S_GAMETICK_DONE_T struct{}
	S_INPUT_T         int
	// Network messages
	N_CONNECT_T     *core.SessionPubConn
	N_DISCONN_T     float64
	N_PLAYER_SYNC_T *PlayerSyncMessage
)

type EventTopic int

func inititializeEventBus() {
	eb = eventBus{}
	eb.subs = make(map[EventTopic]subscribers)

	// Initialize all topics
	for t := EventTopic(0); t < NUMTOPICS; t += 1 {
		eb.subs[t] = subscribers{}
	}
}
