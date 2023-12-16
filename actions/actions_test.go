package actions

import (
	"reflect"
	"regexp"
	"strings"
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

	t.Run("'repite' command returns a RepeatAction instance", func(t *testing.T) {
		// When
		test := Factory("/repite 2 1d20")

		// Assert
		if action := reflect.TypeOf(test).String(); action != "actions.RepeatAction" {
			t.Logf("\nresult expected to be 'actions.RepeatAction'\ngot '%s' instead", action)
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
			t.FailNow()
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
			t.FailNow()
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
			t.FailNow()
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
			t.FailNow()
		}
		if !regexp.MustCompile(`1d20\[\d{1,2}\] = \d{1,2}`).MatchString(test) {
			t.Logf("\nresult expected to be '1d20[<result>] = <total>'\ngot '%s' instead", test)
			t.Fail()
		}
	})

	t.Run("Rolling same dice more than once does not accumulate results", func(t *testing.T) {
		// Given
		action := RollAction{command: "/tira 1d20 + 1d20 +12"}

		// When
		test, err := action.Resolve()

		// Assert
		if err != nil {
			t.Logf("\nresult expected to be 'nil'\ngot '%s' error instead", err.Error())
			t.FailNow()
		}
		if !regexp.MustCompile(`1d20\[\d{1,2}\] \+ 1d20\[\d{1,2}\] \+12 = \d{1,2}`).MatchString(test) {
			t.Logf("\nresult expected to be '1d20[<result>] +1d20[<result>] +12 = <total>'\ngot '%s' instead", test)
			t.Fail()
		}
	})

	t.Run("Rolling negative results deducts on total", func(t *testing.T) {
		// Given
		action := RollAction{command: "/tira 1d20 -1d6 -20"}

		// When
		test, err := action.Resolve()

		// Assert
		if err != nil {
			t.Logf("\nresult expected to be 'nil'\ngot '%s' error instead", err.Error())
			t.FailNow()
		}
		if !regexp.MustCompile(`1d20\[\d{1,2}\] -1d6\[\d{1,2}\] \-20 = -\d{1,2}`).MatchString(test) {
			t.Logf("\nresult expected to be '1d20[<result>] -1d6[<result>] -20 = <total>'\ngot '%s' instead", test)
			t.Fail()
		}
	})
}

func TestRepeatAction(t *testing.T) {
	t.Run("Wrong command returns an MSG_UNKNOWN_ACTION error", func(t *testing.T) {
		// Given
		action := RepeatAction{command: "/repite 10+3d20"}

		// When
		_, test := action.Resolve()

		// Assert
		if test == nil {
			t.Logf("\nresult expected to be an error\ngot 'nil' value instead")
			t.FailNow()
		}
		if test.Error() != MSG_UNKNOWN_ACTION {
			t.Logf("\nresult expected to be '%s'\ngot '%s' instead", MSG_UNKNOWN_ACTION, test.Error())
			t.Fail()
		}
	})

	t.Run("No command to repeat returns an ERR_REPEAT_NODICE error", func(t *testing.T) {
		// Given
		action := RepeatAction{command: "/repite 10 no_dice"}

		// When
		_, test := action.Resolve()

		// Assert
		if test == nil {
			t.Logf("\nresult expected to be an error\ngot 'nil' value instead")
			t.FailNow()
		}
		if test.Error() != ERR_REPEAT_NODICE {
			t.Logf("\nresult expected to be '%s'\ngot '%s' instead", ERR_REPEAT_NODICE, test.Error())
			t.Fail()
		}
	})

	t.Run("Correct command returns correct amount of iterations", func(t *testing.T) {
		// Given
		action := RepeatAction{command: "/repite 10 1d10"}

		// When
		test, err := action.Resolve()

		// Assert
		if err != nil {
			t.Logf("\nresult expected to be nil\ngot '%s' error instead", err.Error())
			t.FailNow()
		}
		if iter := strings.Count(test, "1d10"); iter != 10 {
			t.Logf("\nresult expected to be 10\ngot '%d' instead", iter)
			t.Fail()
		}
	})

	t.Run("Exceeding the maximum number of repetitions returns an ERR_REPEAT_ITER error", func(t *testing.T) {
		// Given
		action := RepeatAction{command: "/repite 11 1d20"}

		// When
		_, test := action.Resolve()

		// Assert
		if test == nil {
			t.Logf("\nresult expected to be an error\ngot 'nil' value instead")
			t.FailNow()
		}
		if test.Error() != ERR_REPEAT_ITER {
			t.Logf("\nresult expected to be '%s'\ngot '%s' instead", ERR_REPEAT_ITER, test.Error())
			t.Fail()
		}
	})

	t.Run("Invalid number of repetitions returns an ERR_REPEAT_ITER error", func(t *testing.T) {
		// Given
		action := RepeatAction{command: "/repite 0 1d10"}

		// When
		_, test := action.Resolve()

		// Assert
		if test == nil {
			t.Logf("\nresult expected to be an error\ngot 'nil' value instead")
			t.FailNow()
		}
		if test.Error() != ERR_REPEAT_ITER {
			t.Logf("\nresult expected to be '%s'\ngot '%s' instead", ERR_REPEAT_ITER, test.Error())
			t.Fail()
		}
	})
}
