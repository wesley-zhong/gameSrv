package player

import (
	"gameSrv/game/gameevent"
	"gameSrv/pkg/event"
	"gameSrv/pkg/log"
)

func InitEvents() {
	event.InitEventDispatcher(1024)
	event.Dispatcher.Register(gameevent.LoginEventID, LoginEventHandler)
	event.Dispatcher.Register(gameevent.DisconnectEventID, DisconnectEventHandler)
}

// event handlers
func LoginEventHandler(event event.Event) {
	log.Infof("on login handler pid =%d eventid %d", event.Player().GetUid(), event.EventId())
}

func DisconnectEventHandler(event event.Event) {
	log.Infof("on disconnected handler pid =%d eventid %d", event.Player().GetUid(), event.EventId())
}
