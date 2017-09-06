package bitmath

import (
	"reflect"
	"testing"
)

var testArr1 = []bool{false, true}
var testArr2 = []bool{false, true, true}
var testArr3 = []bool{true, true, true, true, true, true, true}
var testArr4 = []bool{false}

func TestShift(t *testing.T) {
	actual1 := shift(testArr1, 1)
	expected1 := []bool{false, false, true}

	if !(reflect.DeepEqual(actual1, expected1)) {
		t.Errorf("Test failed: %v != %v\n", actual1, expected1)
	}
}

func TestZero(t *testing.T) {
	actual1 := zero(testArr1)
	expected1 := false

	if actual1 != expected1 {
		t.Errorf("Test failed: %v != %v\n", actual1, expected1)
	}

	actual2 := zero(testArr4)
	expected2 := true

	if actual2 != expected2 {
		t.Errorf("Test failed: %v != %v\n", actual2, expected2)
	}
}

func TestDiv2(t *testing.T) {
	actual1 := div2(testArr1)
	expected1 := []bool{true}

	if !(reflect.DeepEqual(actual1, expected1)) {
		t.Errorf("Test failed: %v != %v\n", actual1, expected1)
	}

	actual2 := div2([]bool{})
	expected2 := []bool{}

	if !(reflect.DeepEqual(actual2, expected2)) {
		t.Errorf("Test failed: %v != %v\n", actual2, expected2)
	}
}

func TestEven(t *testing.T) {
	actual1 := even(testArr1)
	expected1 := true

	if !(reflect.DeepEqual(actual1, expected1)) {
		t.Errorf("Test failed: %v != %v\n", actual1, expected1)
	}

	actual2 := even(testArr3)
	expected2 := false

	if actual2 != expected2 {
		t.Errorf("Test failed: %v != %v\n", actual2, expected2)
	}
}

func TestAdd(t *testing.T) {
	actual1 := add(testArr1, testArr2)
	expected1 := []bool{false, false, false, true}

	if !(reflect.DeepEqual(actual1, expected1)) {
		t.Errorf("Test failed: %v != %v\n", actual1, expected1)
	}

	actual2 := add([]bool{false, true}, add([]bool{false}, []bool{false}))
	expected2 := []bool{false, true}

	if !(reflect.DeepEqual(actual2, expected2)) {
		t.Errorf("Test failed: %v != %v\n", actual2, expected2)
	}
}

func TestExcOr(t *testing.T) {
	actual1 := excOr(false, false, true)
	expected1 := true

	if !(reflect.DeepEqual(actual1, expected1)) {
		t.Errorf("Test failed: %v != %v\n", actual1, expected1)
	}

	actual2 := excOr(false, false, false)
	expected2 := false

	if actual2 != expected2 {
		t.Errorf("Test failed: %v != %v\n", actual2, expected2)
	}
}

func TestNextCarry(t *testing.T) {
	actual1 := nextcarry(false, false, true)
	expected1 := false

	if !(reflect.DeepEqual(actual1, expected1)) {
		t.Errorf("Test failed: %v != %v\n", actual1, expected1)
	}

	actual2 := excOr(false, true, false)
	expected2 := true

	if actual2 != expected2 {
		t.Errorf("Test failed: %v != %v\n", actual2, expected2)
	}
}

func TestMult(t *testing.T) {
	actual1 := mult(testArr4, testArr1)
	expected1 := []bool{false}

	if !(reflect.DeepEqual(actual1, expected1)) {
		t.Errorf("Test failed: %v != %v\n", actual1, expected1)
	}

	actual2 := mult(testArr1, testArr2)
	expected2 := []bool{false, false, true, true}

	if !(reflect.DeepEqual(actual2, expected2)) {
		t.Errorf("Test failed: %v != %v\n", actual2, expected2)
	}
}
