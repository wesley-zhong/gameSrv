package player

import (
	"gameSrv/game/gameevent"
	"gameSrv/game/unlock"
	"gameSrv/pkg/event"
	"gameSrv/pkg/log"
)

func InitEvents() {
	event.InitEventDispatcher(1024)
	event.Dispatcher.Register(gameevent.LoginEventID, LoginEventHandler)
	event.Dispatcher.Register(gameevent.DisconnectEventID, DisconnectEventHandler)

	//
	unlock.InitEvents()
}

// event handlers
func LoginEventHandler(event event.Event) {
	gameEvent := event.(*gameevent.GameEvent)
	log.Infof("on login handler pid =%d eventid %d", gameEvent.PlayerId, event.EventId())
}

func DisconnectEventHandler(event event.Event) {
	gameEvent := event.(*gameevent.GameEvent)
	log.Infof("on disconnected handler pid =%d eventid %d", gameEvent.PlayerId, event.EventId())
}
