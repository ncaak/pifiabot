package actions

import "regexp"

const REGEX_CTH_ACTION = `^/(cth)(\s|$)`
const REGEX_NEWPC_ACTION = `^/(pj)(\s|$)`
const REGEX_REPEAT_ACTION = `^/repite(\s|$)`
const REGEX_ROLL_ACTION = `^/(t|tira)(\s|$)`

const MAX_COMMAND_LENGTH = 100

const ERR_CTH_SKILL_LIMITS = "cth_skill_limits"
const ERR_CTH_WRONG_ARGUMENT = "cth_wrong_argument"
const ERR_UNKNOWN = "unknown_error"
const ERR_REPEAT_ITER = "repeat_iter_error"
const ERR_REPEAT_NODICE = "repeat_nodice_error"

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
	case regexp.MustCompile(REGEX_CTH_ACTION).MatchString(command):
		return CthAction{command: command}
	case regexp.MustCompile(REGEX_NEWPC_ACTION).MatchString(command):
		return NewHeroAction{command: command}
	case regexp.MustCompile(REGEX_REPEAT_ACTION).MatchString(command):
		return RepeatAction{command: command}
	case regexp.MustCompile(REGEX_ROLL_ACTION).MatchString(command):
		return RollAction{command: command}
	default:
		return UnknownAction{}
	}
}
