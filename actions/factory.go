package actions

import "regexp"

const MAX_COMMAND_LENGTH = 50
const MSG_UNKNOWN_ACTION = "unknown_action"
const ERR_UNKNOWN = "unknown_error"

type ActionInterface interface {
	Resolve() (string, error)
}

func Factory(command string) ActionInterface {
	if len(command) > MAX_COMMAND_LENGTH {
		return NewErrorAction("notation_max_length")
	}

	switch {
	// TODO: Define commands as they are coded
	// case strings.HasPrefix(command, "/t"):
	case regexp.MustCompile(`^/(t|tira)(\s|$)`).MatchString(command):
		return RollAction{command: command}
	default:
		return UnknownAction{}
	}
}
