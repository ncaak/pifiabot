package dice

// TODO : Unexport struct and make it work as interface
type Roll struct {
	Symbol string
	Number string
	Faces  string
	Drop   string
}

func NewRoll(symbol string, number string, faces string, drop string) (r Roll) {
	r.Number = number
	r.Faces = faces
	r.Symbol = symbol
	r.Drop = drop

	// No number on the die notation is allowed and it is handled as 1 die [d20 == 1d20]
	if number == "" {
		r.Number = "1"
	}

	return
}
