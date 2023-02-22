package dice

import "testing"

func TestDicePrecheckNumberError(t *testing.T) {

	t.Run("Precheck detects Dice number of 0 and returns 'dice_number' error", func(t *testing.T) {
		// Given
		test := Dice{Number: "0"}

		// When
		err := test.PreCheck()

		// Assert
		if err == nil {
			t.Log("\nresult expected to not be nil")
			t.Fail()

		} else if got := err.Error(); got != "dice_number" {
			t.Log("\nresult expected to have 'dice_number' error \ngot ", got)
			t.Fail()
		}
	})

	t.Run("Precheck detects Dice number greater than 20 and returns 'dice_number' error", func(t *testing.T) {
		// Given
		test := Dice{Number: "21"}

		// When
		err := test.PreCheck()

		// Assert
		if err == nil {
			t.Log("\nresult expected to not be nil")
			t.Fail()

		} else if got := err.Error(); got != "dice_number" {
			t.Log("\nresult expected to have 'dice_number' error \ngot ", got)
			t.Fail()
		}
	})
}

func TestDicePrecheckFacesError(t *testing.T) {

	t.Run("Precheck detects Dice faces of 0 and returns 'faces_number' error", func(t *testing.T) {
		// Given
		test := Dice{Number: "1", Faces: "0"}

		// When
		err := test.PreCheck()

		// Assert
		if err == nil {
			t.Log("\nresult expected to not be nil")
			t.Fail()

		} else if got := err.Error(); got != "faces_number" {
			t.Log("\nresult expected to have 'faces_number' error \ngot ", got)
			t.Fail()
		}
	})

	t.Run("Precheck detects Dice faces greater than 100 and returns 'faces_number' error", func(t *testing.T) {
		// Given
		test := Dice{Number: "1", Faces: "101"}

		// When
		err := test.PreCheck()

		// Assert
		if err == nil {
			t.Log("\nresult expected to not be nil")
			t.Fail()

		} else if got := err.Error(); got != "faces_number" {
			t.Log("\nresult expected to have 'faces_number' error \ngot ", got)
			t.Fail()
		}
	})
}
