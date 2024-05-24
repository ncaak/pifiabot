package actions

import "testing"

func TestCoCActionResolveErrors(t *testing.T) {
	type errorCase struct {
		name    string
		command string
		errType string
	}

	var tests = []errorCase{
		{
			name:    "Failure to send the skill value will return an ERR_COC_WRONG_ARGUMENT error",
			command: "/coc",
			errType: ERR_COC_WRONG_ARGUMENT,
		}, {
			name:    "Sending a non-numeric skill value will return an ERR_COC_WRONG_ARGUMENT error",
			command: "/coc 1d100",
			errType: ERR_COC_WRONG_ARGUMENT,
		}, {
			name:    "Sending a negative skill value will return an ERR_COC_SKILL_LIMITS error",
			command: "/coc -10",
			errType: ERR_COC_SKILL_LIMITS,
		}, {
			name:    "Sending a skill value greater than 100 will return an ERR_COC_SKILL_LIMITS error",
			command: "/coc 101",
			errType: ERR_COC_SKILL_LIMITS,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Given
			action := CoCAction{command: tc.command}

			// When
			_, test := action.Resolve()

			// Assert
			if test == nil {
				t.Logf("\nresult expected to be an error\ngot 'nil' value instead")
				t.FailNow()
			}
			if test.Error() != tc.errType {
				t.Logf("\nresult expected to be '%s'\ngot '%s' instead", tc.errType, test.Error())
				t.Fail()
			}
		})
	}
}

func TestCoCActionResults(t *testing.T) {
	type errorCase struct {
		name       string
		roll       int
		skillValue int
		result     string
	}

	var tests = []errorCase{
		{
			name:       "A roll higher than the skill value will result in a failure",
			roll:       100,
			skillValue: 50,
			result:     "Fallo [100]",
		}, {
			name:       "A roll equal to the skill value will result in a normal success",
			roll:       50,
			skillValue: 50,
			result:     "Éxito (normal) [50]",
		}, {
			name:       "A roll less than the skill value, but not less than half of it, results in a normal success",
			roll:       30,
			skillValue: 50,
			result:     "Éxito (normal) [30]",
		}, {
			name:       "A roll of less than half the skill value results in a hard success",
			roll:       24,
			skillValue: 50,
			result:     "Éxito (difícil) [24]",
		}, {
			name:       "A roll of less than one-fifth of the skill value results in a critical success",
			roll:       10,
			skillValue: 50,
			result:     "Éxito (crítico) [10]",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Given
			action := CoCAction{}

			// When
			test := action.result(tc.skillValue, tc.roll)

			// Assert
			if test != tc.result {
				t.Logf("\nresult expected to be '%s'\ngot '%s' instead", tc.result, test)
				t.Fail()
			}
		})
	}
}
