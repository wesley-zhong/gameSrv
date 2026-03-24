package player

import (
	"gameSrv/pkg/event"
	"gameSrv/pkg/log"
)

const (
	LoginEvent      event.GameEventID = 1
	LogoutEvent     event.GameEventID = 2
	DisconnectEvent event.GameEventID = 3
)

func InitEvents() {
	event.Dispatcher.Register(LoginEvent, LoginEventHandler)
	event.Dispatcher.Register(DisconnectEvent, DisconnectEventHandler)
}

type GameEvent struct {
	player  *GamePlayer
	eventId event.GameEventID
}

func NewGameEvent(player *GamePlayer, eventId event.GameEventID) *GameEvent {
	return &GameEvent{
		player:  player,
		eventId: eventId,
	}
}

func (ge *GameEvent) EventId() event.GameEventID {
	return ge.eventId
}

// event handlers
func LoginEventHandler(event event.Event) {
	gameEvent := event.(*GameEvent)
	log.Infof("on login handler pid =%d eventid %d", gameEvent.player.Id, event.EventId())
}

func DisconnectEventHandler(event event.Event) {
	gameEvent := event.(*GameEvent)
	log.Infof("on disconnected handler pid =%d eventid %d", gameEvent.player.Id, event.EventId())
}
