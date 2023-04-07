package actions

import (
	"fmt"
	"strings"

	"github.com/ncaak/pifiabot/dice"
)

type RollAction struct {
	command string
}

func (a RollAction) Resolve() (string, error) {
	var notation = dice.NewNotation(a.command)
	var message []string
	var total int

	d := notation.GetDice()
	// Command can accept no die to roll 1d20
	if len(d) == 0 {
		d = append(d, dice.D20())
	}

	// Append dice algebra and their results e.g. 1d20[20]
	for _, die := range d {
		if err := die.PreCheck(); err != nil {
			return "", err
		}

		results, subtotal := die.Roll()

		message = append(message, fmt.Sprintf("%s%v", die.GetAlgebra(), results))
		total += subtotal
	}

	// Append bonuses in case there are any
	if bonusStr, bonusTotal := notation.GetBonusAndTotal(); bonusStr != "" {
		message = append(message, bonusStr)
		total += bonusTotal
	}

	// Append total before sending the message
	message = append(message, fmt.Sprintf("= %d", total))

	return strings.Join(message, " "), nil
}

// TODO> drops
