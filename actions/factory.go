package actions

import (
	"strings"
)

const MAX_COMMAND_LENGTH = 50

type Action interface {
	Resolve() (string, error)
}

func Factory(command string) Action {
	switch {
	// TODO: Define commands as they are coded
	case strings.HasPrefix(command, "/t"):
		return RollAction{command: command}
	default:
		return UnknownAction{}
	}
}
