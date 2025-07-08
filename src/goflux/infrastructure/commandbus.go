package infrastructure

import (
	"goflux/application"
)

type CommandBus interface {
	// Dispatch sends a command to the appropriate handler.
	Dispatch(command application.Command) error
	// RegisterHandler registers a command handler for a specific command type.
	RegisterHandler(commandType string, handler application.CommandHandler) error
	// UnregisterHandler removes a command handler for a specific command type.
}
