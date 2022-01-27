package inmemory

import (
	"bunsan-ocr/kit/bus/event"
	"context"
	"fmt"
	"sync"
)

// EventBus is an in-memory implementation of the event.Bus.
type EventBus struct {
	handlers map[event.Type][]event.Handler
	wg       sync.WaitGroup
}

// NewEventBus initializes a new EventBus.
func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[event.Type][]event.Handler),
		wg:       sync.WaitGroup{},
	}
}

// Subscribe implements the event.Bus interface.
func (b *EventBus) Subscribe(evtType event.Type, handler event.Handler) {
	subscribersForType, ok := b.handlers[evtType]
	if !ok {
		b.handlers[evtType] = []event.Handler{handler}
	}

	subscribersForType = append(subscribersForType, handler)
}

// Publish implements the event.Bus interface.
func (b *EventBus) Publish(ctx context.Context, events []event.Event) error {
	for _, evt := range events {
		handlers, ok := b.handlers[evt.Type()]
		if !ok {
			return nil
		}

		for _, handler := range handlers {
			b.wg.Add(1)
			go b.doPublish(ctx, evt, handler)
		}
	}

	return nil
}

func (b *EventBus) doPublish(ctx context.Context, evt event.Event, handler event.Handler) {
	defer b.wg.Done()
	if err := handler.Handle(ctx, evt); err != nil {
		fmt.Println(handler.Type(), err)
	}
}
