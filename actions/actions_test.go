package actions

import "testing"

func TestRollAction(t *testing.T) {
	var action = RollAction{}

	t.Run("Resolve returns 'unknown_action' error if command is not correct", func(t *testing.T) {
		// Given
		action.command = "/tirar"
		expectedError := "unknown_action"

		// When
		_, result := action.Resolve()

		// Assert
		if result == nil || result.Error() != expectedError {
			t.Logf("\nresult should be '%s' error\ngot %s", expectedError, result.Error())
			t.Fail()
		}
	})

	t.Run("Resolve returns 'notation_max_length' error if notation length goes over the limit", func(t *testing.T) {
		// Given
		action.command = "/tira 000000000000000000000000000000000000000000000"
		expectedError := "notation_max_length"

		// When
		_, result := action.Resolve()

		// Assert
		if result == nil || result.Error() != expectedError {
			t.Logf("\nresult should be '%s' error\ngot %s", expectedError, result.Error())
			t.Fail()
		}
	})
}
