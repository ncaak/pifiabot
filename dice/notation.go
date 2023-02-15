package dice

import (
	"regexp"
	"strings"
)

const REGEX_DICE_NOTATION = `([-\+\s])(\d*)d(\d+)(-[HL])?`
const REGEX_BONUS_NOTATION = `([+-]\d+)`

type Notation struct {
	dice  []Dice
	text  string
	Bonus string
}

func NewNotation(message string) (n Notation) {
	n.text = message
	n.dice = n.extractDice()

	return
}

func (n Notation) extractBonus() []string {
	var bonus []string
	var re = regexp.MustCompile(REGEX_BONUS_NOTATION)

	for _, match := range re.FindAllStringSubmatch(n.noDiceText(), -1) {
		bonus = append(bonus, match[1])
	}

	return bonus
}

func (n Notation) extractDice() []Dice {
	var diceRoll []Dice
	var re = regexp.MustCompile(REGEX_DICE_NOTATION)

	for _, match := range re.FindAllStringSubmatch(n.text, -1) {
		diceRoll = append(diceRoll, Dice{
			Algebra: match[0],
			Symbol:  match[1],
			Number:  match[2],
			Faces:   match[3],
			Drop:    match[4],
		})
	}

	return diceRoll
}

func (n Notation) noDiceText() string {
	var t = n.text
	for i := range n.dice {
		t = strings.Replace(t, n.dice[i].Algebra, "", 1)
	}
	t = strings.ReplaceAll(t, " ", "")
	return t
}
