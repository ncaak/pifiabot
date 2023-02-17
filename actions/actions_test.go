package actions

import (
	"reflect"
	"testing"
)

func TestActionsFactory(t *testing.T) {
	t.Run("unknown command returns an UnknownAction instance", func(t *testing.T) {
		// When
		test := Factory("/mock")

		// Assert
		if action := reflect.TypeOf(test).String(); action != "actions.UnknownAction" {
			t.Logf("\nresult expected to be 'actions.UnknownAction'\ngot '%s' instead", action)
			t.Fail()
		}
	})

	t.Run("command length is over the limit returns an ErrorAction instance", func(t *testing.T) {
		// When
		test := Factory("/tira 000000000000000000000000000000000000000000000")

		// Assert
		if action := reflect.TypeOf(test).String(); action != "actions.ErrorAction" {
			t.Logf("\nresult expected to be 'actions.ErrorAction'\ngot '%s' instead", action)
			t.Fail()
		}
	})

	t.Run("'tira' command returns a RollAction instance", func(t *testing.T) {
		// When
		test := Factory("/tira 1d20")

		// Assert
		if action := reflect.TypeOf(test).String(); action != "actions.RollAction" {
			t.Logf("\nresult expected to be 'actions.RollAction'\ngot '%s' instead", action)
			t.Fail()
		}
	})

	t.Run("'t' command returns a RollAction instance", func(t *testing.T) {
		// When
		test := Factory("/t 1d20")

		// Assert
		if action := reflect.TypeOf(test).String(); action != "actions.RollAction" {
			t.Logf("\nresult expected to be 'actions.RollAction'\ngot '%s' instead", action)
			t.Fail()
		}
	})
}

func TestRollAction(t *testing.T) {

}
