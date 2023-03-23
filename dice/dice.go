package dice

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const MAX_DICE_NUMER = 20
const MAX_DICE_FACES = 100

// TODO : Unexport struct and make it work as interface
type Dice struct {
	Algebra string
	Symbol  string
	Number  string
	Faces   string
	Drop    string
	number  int
	faces   int
}

func (d *Dice) PreCheck() error {
	var err error

	if d.Number == "" {
		d.number = 1

	} else {
		d.number, err = strconv.Atoi(d.Number)
		if err != nil || d.number < 1 || d.number > MAX_DICE_NUMER {
			return fmt.Errorf("dice_number")
		}
	}

	d.faces, err = strconv.Atoi(d.Faces)
	if err != nil || d.faces < 1 || d.faces > MAX_DICE_FACES {
		return fmt.Errorf("faces_number")
	}

	return err
}

func (d Dice) Roll() (results []int, total int) {
	rand.Seed(time.Now().UnixNano())

	i := 0
	for i < d.number {
		r := rand.Intn(d.faces) + 1
		results = append(results, r)
		total += r
		i++
	}

	return
}
