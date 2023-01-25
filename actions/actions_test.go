package actions

import "testing"

func TestBaseAction(t *testing.T) {
	var baseAction = baseAction{}

	t.Run("extractDice two items on the notation", func(t *testing.T) {
		// Given
		roll := "/command 1d20+2d4"

		// When
		result := baseAction.extractDice(roll)

		// Assert
		if got := len(result); got != 2 {
			t.Log("\nresult expected to have 2 items\ngot ", got)
			t.Fail()
		}
	})

	t.Run("extractDice accepts no die number (meaning 1 die)", func(t *testing.T) {
		// Given
		roll := "/command d20"

		// When
		result := baseAction.extractDice(roll)

		// Assert
		if got := len(result); got != 1 {
			t.Log("\nresult expected to have 1 item\ngot ", got)
			t.Fail()

		} else if r := result[0]; r.Number != "1" || r.Faces != "20" {
			t.Logf("\nresult expected to have 1 die of 20 faces\ngot %s dice of %s faces", r.Number, r.Faces)
			t.Fail()
		}
	})

	t.Run("extractDice misspell notation", func(t *testing.T) {
		// Given
		roll := "/command 1d20d8"

		// When
		result := baseAction.extractDice(roll)

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
		roll := "/command 1d20-1d8"

		// When
		result := baseAction.extractDice(roll)

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
		roll := "/command 4d6-L"

		// When
		result := baseAction.extractDice(roll)

		// Assert
		if got := len(result); got != 1 {
			t.Log("\nresult expected to have 1 item\ngot ", got)
			t.Fail()

		} else if r := result[0]; r.Drop != "-L" {
			t.Logf("\nresult expected to have -L drop suffix\ngot %s suffix", r.Drop)
			t.Fail()
		}
	})
}

func TestRollAction(t *testing.T) {
	var action = RollAction{}

	t.Run("extractDice get an error if notation goes over the limit", func(t *testing.T) {
		// Given
		action.command = "/tira 000000000000000000000000000000000000000000000"

		// When
		_, result := action.Resolve()

		// Assert
		if result == nil || result.Error() != "notation_max_length" {
			t.Logf("\nresult should be 'notation_max_length' error\ngot %s", result.Error())
			t.Fail()
		}
	})
}
