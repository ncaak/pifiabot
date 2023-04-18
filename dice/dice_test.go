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

func TestPrecheckDrop(t *testing.T) {
	t.Run("Precheck detects drops with only one die and returns 'no_drop' error", func(t *testing.T) {
		// Given
		test := Dice{Number: "1", Faces: "20", Drop: "L"}

		// When
		err := test.PreCheck()

		// Assert
		if err == nil {
			t.Log("\nresult expected to not be nil")
			t.Fail()

		} else if got := err.Error(); got != "no_drop" {
			t.Log("\nresult expected to have 'no_drop' error \ngot ", got)
			t.Fail()
		}
	})
}

func TestDropFromTotal(t *testing.T) {
	t.Run("dropFromTotal returns minimum value from slice", func(t *testing.T) {
		// Given
		test := Dice{facesVal: 20, Drop: "-L"}
		mock := []int{7, 10, 20, 3, 8}

		// When
		result := test.dropFromTotal(mock)

		// Assert
		if result != 3 {
			t.Logf("\nresult expected to be 3\ngot %d instead", result)
			t.Fail()
		}
	})

	t.Run("dropFromTotal returns maximum value from slice", func(t *testing.T) {
		// Given
		test := Dice{facesVal: 20, Drop: "-H"}
		mock := []int{7, 10, 20, 3, 8}

		// When
		result := test.dropFromTotal(mock)

		// Assert
		if result != 20 {
			t.Logf("\nresult expected to be 20\ngot %d instead", result)
			t.Fail()
		}
	})

	t.Run("dropFromTotal returns 0 if dropping flag is unknown", func(t *testing.T) {
		// Given
		test := Dice{facesVal: 20, Drop: "unknown"}
		mock := []int{7, 10, 20, 3, 8}

		// When
		result := test.dropFromTotal(mock)

		// Assert
		if result != 0 {
			t.Logf("\nresult expected to be 0\ngot %d instead", result)
			t.Fail()
		}
	})
}
