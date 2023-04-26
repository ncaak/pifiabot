package actions

import "regexp"

const REGEX_ROLL_ACTION = `^/(t|tira)(\s|$)`
const REGEX_REPEAT_ACTION = `^/repite(\s|$)`

const MAX_COMMAND_LENGTH = 50
const ERR_UNKNOWN = "unknown_error"

const MSG_UNKNOWN_ACTION = "unknown_action"

type ActionInterface interface {
	Resolve() (string, error)
}

func Factory(command string) ActionInterface {
	if len(command) > MAX_COMMAND_LENGTH {
		return NewErrorAction("notation_max_length")
	}

	switch {
	// TODO: Define commands as they are coded
	case regexp.MustCompile(REGEX_ROLL_ACTION).MatchString(command):
		return RollAction{command: command}
	case regexp.MustCompile(REGEX_REPEAT_ACTION).MatchString(command):
		return RepeatAction{command: command}
	default:
		return UnknownAction{}
	}
}
