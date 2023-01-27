package actions

import (
	"regexp"
	"strings"

	"github.com/ncaak/pifiabot/dice"
)

const MAX_COMMAND_LENGTH = 50
const REGEX_DICE_NOTATION = `([-\+\s])(\d*)d(\d+)(-[HL])?`
const REGEX_BONUS_NOTATION = `([+-]\d+)`

type baseAction struct{}

func (a baseAction) extractBonus(notation string) []string {
	var bonus []string
	var re = regexp.MustCompile(REGEX_BONUS_NOTATION)
	// Removes first dice notation then it trims all whitespaces
	var nakedNotation = strings.ReplaceAll(
		regexp.MustCompile(REGEX_DICE_NOTATION).ReplaceAllString(notation, ""),
		" ", "",
	)

	for _, match := range re.FindAllStringSubmatch(nakedNotation, -1) {
		bonus = append(bonus, match[1])
	}

	return bonus
}

func (a baseAction) extractDice(notation string) []dice.Dice {
	var diceRoll []dice.Dice
	var re = regexp.MustCompile(REGEX_DICE_NOTATION)

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
