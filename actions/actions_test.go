package actions

import (
	"reflect"
	"regexp"
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
	t.Run("Rolling a simple roll (tira) returns a good formatted message", func(t *testing.T) {
		// Given
		action := RollAction{command: "/tira 1d20"}

		// When
		test, err := action.Resolve()

		// Assert
		if err != nil {
			t.Logf("\nresult expected to be 'nil'\ngot '%s' error instead", err.Error())
			t.Fail()
		}
		if !regexp.MustCompile(`1d20\[\d{1,2}\] = \d{1,2}`).MatchString(test) {
			t.Logf("\nresult expected to be '1d20[<result>] = <total>'\ngot '%s' instead", test)
			t.Fail()
		}
	})

	t.Run("Rolling a simple roll (t) returns a good formatted message", func(t *testing.T) {
		// Given
		action := RollAction{command: "/t 1d20"}

		// When
		test, err := action.Resolve()

		// Assert
		if err != nil {
			t.Logf("\nresult expected to be 'nil'\ngot '%s' error instead", err.Error())
			t.Fail()
		}
		if !regexp.MustCompile(`1d20\[\d{1,2}\] = \d{1,2}`).MatchString(test) {
			t.Logf("\nresult expected to be '1d20[<result>] = <total>'\ngot '%s' instead", test)
			t.Fail()
		}
	})

	t.Run("Rolling a simple roll (tira) returns 1d20 by default", func(t *testing.T) {
		// Given
		action := RollAction{command: "/tira"}

		// When
		test, err := action.Resolve()

		// Assert
		if err != nil {
			t.Logf("\nresult expected to be 'nil'\ngot '%s' error instead", err.Error())
			t.Fail()
		}
		if !regexp.MustCompile(`1d20\[\d{1,2}\] = \d{1,2}`).MatchString(test) {
			t.Logf("\nresult expected to be '1d20[<result>] = <total>'\ngot '%s' instead", test)
			t.Fail()
		}
	})

	t.Run("Rolling a simple roll (t) returns 1d20 by default", func(t *testing.T) {
		// Given
		action := RollAction{command: "/t"}

		// When
		test, err := action.Resolve()

		// Assert
		if err != nil {
			t.Logf("\nresult expected to be 'nil'\ngot '%s' error instead", err.Error())
			t.Fail()
		}
		if !regexp.MustCompile(`1d20\[\d{1,2}\] = \d{1,2}`).MatchString(test) {
			t.Logf("\nresult expected to be '1d20[<result>] = <total>'\ngot '%s' instead", test)
			t.Fail()
		}
	})
}
