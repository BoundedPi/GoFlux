package application

import (
	"goflux/domain"
)

type EventHandler interface {
	GetEventIdentifier() string
	Handle(event domain.Event) error
}
