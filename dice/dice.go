package dice

import (
	"fmt"
	"math/rand"
	"strconv"
)

const MAX_DICE_NUMBER = 20
const MAX_DICE_FACES = 100

type Dice struct {
	Algebra   string
	Symbol    string
	Number    string
	Faces     string
	Drop      string
	numberVal int
	facesVal  int
}

func (d Dice) GetAlgebra() string {
	return d.Algebra
}

func (d *Dice) PreCheck() error {
	var err error

	if d.Number == "" {
		d.numberVal = 1

	} else {
		d.numberVal, err = strconv.Atoi(d.Number)
		if err != nil || d.numberVal < 1 || d.numberVal > MAX_DICE_NUMBER {
			return fmt.Errorf("dice_number")
		}
	}

	d.facesVal, err = strconv.Atoi(d.Faces)
	if err != nil || d.facesVal < 1 || d.facesVal > MAX_DICE_FACES {
		return fmt.Errorf("faces_number")
	}

	if d.Drop != "" && d.numberVal == 1 {
		return fmt.Errorf("no_drop")
	}

	return err
}

func (d Dice) Roll() (results []int, total int) {
	i := 0
	for i < d.numberVal {
		r := rand.Intn(d.facesVal) + 1
		results = append(results, r)

		if d.Symbol == "-" {
			r *= -1
		}
		total += r

		i++
	}

	if d.Drop != "" {
		total -= d.dropFromTotal(results)
	}

	return
}

func (d Dice) dropFromTotal(results []int) int {
	var initValue int
	var from func(int, int) int

	switch d.Drop {
	case "-L":
		initValue = d.facesVal
		from = func(c int, v int) int {
			if c > v {
				return v
			}
			return c
		}

	case "-H":
		initValue = 0
		from = func(c int, v int) int {
			if c < v {
				return v
			}
			return c
		}
	default:
		return 0
	}

	valueToDrop := initValue
	for _, r := range results {
		valueToDrop = from(valueToDrop, r)
	}

	return valueToDrop
}
