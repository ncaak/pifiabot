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
	var results = make(map[string][]int)
	var total int

	d := notation.GetDice()
	// Command can accept no die to roll 1d20
	if len(d) == 0 {
		d = append(d, dice.D20())
	}

	for _, dice := range d {
		if err := dice.PreCheck(); err != nil {
			return "", err
		}

		subtotal := 0
		results[dice.Algebra], subtotal = dice.Roll()
		total += subtotal
	}

	// TODO Bonus

	return a.format(results, total), nil
}

func (a RollAction) format(results map[string][]int, total int) string {
	var msg []string
	for k, v := range results {
		msg = append(msg, fmt.Sprintf("%s%v", k, v))
	}

	msg = append(msg, fmt.Sprintf("= %d", total))

	return strings.Join(msg, " ")
}
