package dice

func D20() Dice {
	return Dice{Algebra: "1d20", Faces: "20"}
}

func RollD100() int {
	d := Dice{numberVal: 1, facesVal: 100}
	_, t := d.Roll()
	return t
}
