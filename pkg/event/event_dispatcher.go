package event

type GameEventID int32
type Event interface {
	EventId() GameEventID
}

type Handler func(event Event)
type DispatcherWrapper struct {
	eventHandlers map[GameEventID][]Handler
}

var Dispatcher *DispatcherWrapper

func InitEventDispatcher(count int32) {
	Dispatcher = newEventDispatcher(count)
}

func newEventDispatcher(count int32) *DispatcherWrapper {
	return &DispatcherWrapper{
		eventHandlers: make(map[GameEventID][]Handler, count),
	}
}

func (ed *DispatcherWrapper) Register(eventId GameEventID, handler Handler) {
	ed.eventHandlers[eventId] = append(ed.eventHandlers[eventId], handler)
}

func (ed *DispatcherWrapper) Dispatch(event Event) {
	for _, handler := range ed.eventHandlers[event.EventId()] {
		handler(event)
	}
}
