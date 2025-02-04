package aggregate

import (
	"ddd-sample/internal/core/event"
	"sync"
)

type CoreAggregate interface {
	AddEvent(e event.Event)
	PopEvents() []event.Event
}
type coreAggregate struct {
	events []event.Event
	mutex  *sync.Mutex
}

// 工廠
func NewCoreAggregate() CoreAggregate {
	return &coreAggregate{
		events: make([]event.Event, 0),
		mutex:  new(sync.Mutex),
	}
}

func (ca *coreAggregate) AddEvent(e event.Event) {
	ca.mutex.Lock()
	defer ca.mutex.Unlock()

	ca.events = append(ca.events, e)
}

func (ca *coreAggregate) PopEvents() []event.Event {
	ca.mutex.Lock()
	defer ca.mutex.Unlock()

	return ca.events
}
