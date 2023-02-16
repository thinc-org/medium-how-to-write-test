package example

import "testing"

func TestDivideSuccess(t *testing.T) {
	want := 10
	input1 := 50
	input2 := 5

	actual, err := Divide(input1, input2)

	if err != nil {
		t.Errorf("got an error %s", err.Error())
		return
	}

	if want != actual {
		t.Errorf("actual: %v not equal to expected: %v", actual, want)
		return
	}
}

func TestDivideByZeroError(t *testing.T) {
	want := 0
	input1 := 50
	input2 := 0

	actual, err := Divide(input1, input2)

	if err == nil {
		t.Errorf("should got error when divide by zero")
		return
	}

	if actual != want {
		t.Errorf("should be default value (0)")
		return
	}
}
