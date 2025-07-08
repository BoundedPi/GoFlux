package infrastructure

import (
	"goflux/application"
	"goflux/domain"
)

type EventBus interface {
	// Publish publishes an event to the event bus.
	Publish(event domain.Event) error
	// Subscribe subscribes a handler to a specific event type.
	Subscribe(eventType string, handler application.EventHandler) error
	// Unsubscribe removes a handler from a specific event type.
	Unsubscribe(eventType string, handler application.EventHandler) error
}
