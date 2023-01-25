package dice

// TODO : Unexport struct and make it work as interface
type Dice struct {
	Symbol string
	Number string
	Faces  string
	Drop   string
}

func NewDice(symbol string, number string, faces string, drop string) (d Dice) {
	d.Number = number
	d.Faces = faces
	d.Symbol = symbol
	d.Drop = drop

	// No number on the die notation is allowed and it is handled as 1 die [d20 == 1d20]
	if number == "" {
		d.Number = "1"
	}

	return
}
