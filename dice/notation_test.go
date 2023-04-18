package dice

import "testing"

func TestNotation(t *testing.T) {

	t.Run("extractDice two items on the notation", func(t *testing.T) {
		// Given
		var notation = NewNotation("/command 1d20+2d4")

		// When
		result := notation.extractDice()

		// Assert
		if got := len(result); got != 2 {
			t.Log("\nresult expected to have 2 items\ngot ", got)
			t.Fail()
		}
	})

	t.Run("extractDice accepts no die number (meaning 1 die)", func(t *testing.T) {
		// Given
		var notation = NewNotation("/command d20")

		// When
		result := notation.extractDice()

		// Assert
		if got := len(result); got != 1 {
			t.Log("\nresult expected to have 1 item\ngot ", got)
			t.Fail()

		} else if r := result[0]; r.Number != "" || r.Faces != "20" {
			t.Logf("\nresult expected to have a die of 20 faces\ngot %s dice of %s faces", r.Number, r.Faces)
			t.Fail()
		}
	})

	t.Run("extractDice misspell notation", func(t *testing.T) {
		// Given
		var notation = NewNotation("/command 1d20d8")

		// When
		result := notation.extractDice()

		// Assert
		if got := len(result); got != 1 {
			t.Log("\nresult expected to have 1 item\ngot ", got)
			t.Fail()

		} else if r := result[0]; r.Number != "1" || r.Faces != "20" {
			t.Logf("\nresult expected to have 1 die of 20 faces\ngot %s dice of %s faces", r.Number, r.Faces)
			t.Fail()
		}
	})

	t.Run("extractDice substraction symbol is catched", func(t *testing.T) {
		// Given
		var notation = NewNotation("/command 1d20-1d8")

		// When
		result := notation.extractDice()

		// Assert
		if got := len(result); got != 2 {
			t.Log("\nresult expected to have 2 item\ngot ", got)
			t.Fail()

		} else if r := result[1]; r.Symbol != "-" {
			t.Logf("\nresult expected to have - symbol\ngot (%s) symbol", r.Symbol)
			t.Fail()
		}
	})

	t.Run("extractDice notation with drop suffix", func(t *testing.T) {
		// Given
		var notation = NewNotation("/command 4d6-L")

		// When
		result := notation.extractDice()

		// Assert
		if got := len(result); got != 1 {
			t.Log("\nresult expected to have 1 item\ngot ", got)
			t.Fail()

		} else if r := result[0]; r.Drop != "-L" {
			t.Logf("\nresult expected to have -L drop suffix\ngot %s suffix", r.Drop)
			t.Fail()
		}
	})

	t.Run("extractBonus does not match on multiple dice notation", func(t *testing.T) {
		// Given
		var notation = NewNotation("/command 1d20+6d6")

		// When
		result := notation.extractBonus()

		// Assert
		if got := len(result); got != 0 {
			t.Log("\nresult expected to have 0 item\ngot ", got)
			t.Fail()
		}
	})

	t.Run("extractBonus multiple bonus", func(t *testing.T) {
		// Given
		var notation = NewNotation("/command 2d6+2 + 7")

		// When
		result := notation.extractBonus()

		// Assert
		if got := len(result); got != 2 {
			t.Log("\nresult expected to have 2 item\ngot ", got)
			t.Fail()

		} else if r := result[1]; r != "+7" {
			t.Logf("\nresult expected to have '7'\ngot %s", r)
			t.Fail()
		}
	})

	t.Run("extractBonus negative and positive bonus mixed", func(t *testing.T) {
		// Given
		var notation = NewNotation("/command 2d6+2 -7")

		// When
		result := notation.extractBonus()

		// Assert
		if got := len(result); got != 2 {
			t.Log("\nresult expected to have 2 item\ngot ", got)
			t.Fail()

		} else if r := result[1]; r != "-7" {
			t.Logf("\nresult expected to have '-7'\ngot %s", r)
			t.Fail()
		}
	})

}
