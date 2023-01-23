package actions

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ncaak/pifiabot/dice"
)

const MAX_COMMAND_LENGTH = 50

type baseAction struct{}

func (a baseAction) extractDice(notation string) (rolls []dice.Roll, err error) { //TODO: return error on prechecks
	if len(notation) > MAX_COMMAND_LENGTH {
		err = fmt.Errorf("notation_max_length")
		return
	}

	var re = regexp.MustCompile(`([-\+\s])(\d*)d(\d+)(-[HL])?`)
	var matches = re.FindAllStringSubmatch(notation, -1)

	for _, match := range matches {
		rolls = append(rolls, dice.NewRoll(
			match[1],
			match[2],
			match[3],
			match[4],
		))
	}

	return
}

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
