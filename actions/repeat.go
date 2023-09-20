package actions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ncaak/pifiabot/dice"
)

type RepeatAction struct {
	command    string
	iterations int
	notation   dice.Notation
}

func (a RepeatAction) Resolve() (string, error) {
	var err error
	var params []string = strings.SplitN(a.command, " ", 3)
	// Get reps or return an error
	a.iterations, err = strconv.Atoi(params[1])
	if err != nil {
		return "", fmt.Errorf(MSG_UNKNOWN_ACTION)
	}

	// Get notation that will be the same for every rep
	a.notation, err = a.getNotation(params[2])
	if err != nil {
		return "", err
	}

	var results []string
	for i := 0; i < a.iterations; i++ {
		results = append(results, a.solveNotation(a.notation))
	}

	return strings.Join(results, "\n"), nil
}

func (a RepeatAction) getNotation(algebra string) (dice.Notation, error) {
	var notation = dice.NewNotation(algebra)
	d := notation.GetDice()
	if len(d) == 0 {
		return notation, fmt.Errorf(ERR_REPEAT)
	}

	for _, die := range d {
		err := die.PreCheck()
		if err != nil {
			return notation, err
		}
	}

	return notation, nil
}

func (a RepeatAction) solveNotation(n dice.Notation) string {
	var message []string
	var total int

	d := n.GetDice()
	// Command can accept no die to roll 1d20
	if len(d) == 0 {
		d = append(d, dice.D20())
	}

	// Append dice algebra and their results e.g. 1d20[20]
	for _, die := range d {
		results, subtotal := die.Roll()

		message = append(message, fmt.Sprintf("%s%v", die.GetAlgebra(), results))
		total += subtotal
	}

	// Append bonuses in case there are any
	if bonusStr, bonusTotal := n.GetBonusAndTotal(); bonusStr != "" {
		message = append(message, bonusStr)
		total += bonusTotal
	}

	// Append total before sending the message
	message = append(message, fmt.Sprintf("= %d", total))

	return strings.Join(message, " ")
}
