package actions

import (
	"regexp"
	"strings"

	"github.com/ncaak/pifiabot/dice"
)

const MAX_COMMAND_LENGTH = 50

type baseAction struct{}

func (a baseAction) extractDice(notation string) []dice.Dice {
	var diceRoll []dice.Dice
	var re = regexp.MustCompile(`([-\+\s])(\d*)d(\d+)(-[HL])?`)

	for _, match := range re.FindAllStringSubmatch(notation, -1) {
		diceRoll = append(diceRoll, dice.NewDice(
			match[1],
			match[2],
			match[3],
			match[4],
		))
	}

	return diceRoll
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
