package application

import (
	"goflux/domain/values"
)

type CommandHandler interface {
	GetCommandIdentifier() values.CommandId
	Handle(command Command) error
}
