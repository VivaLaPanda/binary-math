package bitmath

import (
	"math"
	"reflect"
	"testing"
)

var testArr1 = []bool{false, true}
var testArr2 = []bool{false, true, true}
var testArr3 = []bool{true, true, true, true, true, true, true}
var testArr4 = []bool{false}

func TestIntMult(t *testing.T) {
	actual1 := IntMult(2, 3)
	expected1 := 6

	if !(reflect.DeepEqual(actual1, expected1)) {
		t.Errorf("Test failed: %v != %v\n", actual1, expected1)
	}

	actual2 := IntMult(100, 100)
	expected2 := 10000

	if !(reflect.DeepEqual(actual2, expected2)) {
		t.Errorf("Test failed: %v != %v\n", actual2, expected2)
	}
}

func TestIntAdd(t *testing.T) {
	actual1 := IntAdd(2, 3)
	expected1 := 5

	if !(reflect.DeepEqual(actual1, expected1)) {
		t.Errorf("Test failed: %v != %v\n", actual1, expected1)
	}

	actual2 := IntAdd(math.MaxInt64/3, math.MaxInt64/3)
	expected2 := (math.MaxInt64 / 3) + (math.MaxInt64 / 3)

	if !(reflect.DeepEqual(actual2, expected2)) {
		t.Errorf("Test failed: %v != %v\n", actual2, expected2)
	}
}

func TestIntSub(t *testing.T) {
	actual1 := IntSub(5, 1)
	expected1 := 4

	if !(reflect.DeepEqual(actual1, expected1)) {
		t.Errorf("Test failed: %v != %v\n", actual1, expected1)
	}

	actual2 := IntSub(1, 5)
	expected2 := -4

	if actual2 != expected2 {
		t.Errorf("Test failed: %v != %v\n", actual2, expected2)
	}

	actual3 := IntSub(math.MaxInt64, 1)
	expected3 := math.MaxInt64 - 1

	if actual3 != expected3 {
		t.Errorf("Test failed: %v != %v\n", actual3, expected3)
	}
}

func TestIntExp(t *testing.T) {
	actual1 := IntExp(2, 3)
	expected1 := 8

	if !(reflect.DeepEqual(actual1, expected1)) {
		t.Errorf("Test failed: %v != %v\n", actual1, expected1)
	}

	actual2 := IntExp(10, 11)
	expected2 := 100000000000

	if actual2 != expected2 {
		t.Errorf("Test failed: %v != %v\n", actual2, expected2)
	}
}

func TestIntDivide(t *testing.T) {
	actual1q, actual1r := IntDivide(10, 2)
	expected1q, expected1r := 5, 0

	if !(reflect.DeepEqual(actual1q, expected1q)) {
		t.Errorf("Test failed: %v != %v\n", actual1q, expected1q)
	}

	if !(reflect.DeepEqual(actual1r, expected1r)) {
		t.Errorf("Test failed: %v != %v\n", actual1r, expected1r)
	}
}

func TestShift(t *testing.T) {
	actual1 := shift(testArr1, 1)
	expected1 := []bool{false, false, true}

	if !(reflect.DeepEqual(actual1, expected1)) {
		t.Errorf("Test failed: %v != %v\n", actual1, expected1)
	}
}

func TestNbitPrime(t *testing.T) {
	actual1 := Bin2dec(NBitPrime(Dec2bin(8), Dec2bin(10)))
	expected1 := 149

	if actual1 != expected1 {
		t.Errorf("Test failed: %v != %v\n", actual1, expected1)
	}

	actual2 := Bin2dec(NBitPrime(Dec2bin(10), Dec2bin(10)))
	expected2 := 617

	if actual2 != expected2 {
		t.Errorf("Test failed: %v != %v\n", actual2, expected2)
	}
}

func TestPrimality(t *testing.T) {
	actual1 := PrimalityThree(Dec2bin(100), Dec2bin(10))
	expected1 := false

	if actual1 != expected1 {
		t.Errorf("Test failed: %v != %v\n", actual1, expected1)
	}

	actual2 := PrimalityThree(Dec2bin(23), Dec2bin(20))
	expected2 := true

	if actual2 != expected2 {
		t.Errorf("Test failed: %v != %v\n", actual2, expected2)
	}

	actual3 := PrimalityThree(Dec2bin(437), Dec2bin(10))
	expected3 := false

	if actual3 != expected3 {
		t.Errorf("Test failed: %v != %v\n", actual3, expected3)
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
	actual1 := Add(testArr1, testArr2)
	expected1 := []bool{false, false, false, true}

	if !(reflect.DeepEqual(actual1, expected1)) {
		t.Errorf("Test failed: %v != %v\n", actual1, expected1)
	}

	actual2 := Add([]bool{false, true}, Add([]bool{false}, []bool{false}))
	expected2 := []bool{false, true}

	if !(reflect.DeepEqual(actual2, expected2)) {
		t.Errorf("Test failed: %v != %v\n", actual2, expected2)
	}
}

func TestSub(t *testing.T) {
	actual1, true := Sub(testArr1, testArr2)
	expected1 := []bool{false, false, true}

	if !(reflect.DeepEqual(actual1, expected1)) {
		t.Errorf("Test failed: %v != %v\n", actual1, expected1)
	}

	actual2, false := Sub(testArr3, testArr1)
	expected2 := []bool{true, false, true, true, true, true, true}

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
	actual1 := Mult(testArr4, testArr1)
	expected1 := []bool{false}

	if !(reflect.DeepEqual(actual1, expected1)) {
		t.Errorf("Test failed: %v != %v\n", actual1, expected1)
	}

	actual2 := Mult(testArr1, testArr2)
	expected2 := []bool{false, false, true, true}

	if !(reflect.DeepEqual(actual2, expected2)) {
		t.Errorf("Test failed: %v != %v\n", actual2, expected2)
	}
}

func TestDec2Bin(t *testing.T) {
	actual1 := Dec2bin(4)
	expected1 := []bool{false, false, true, false}

	if !(reflect.DeepEqual(actual1, expected1)) {
		t.Errorf("Test failed: %v != %v\n", actual1, expected1)
	}

	actual2 := Dec2bin(0)
	expected2 := []bool{false}

	if !(reflect.DeepEqual(actual2, expected2)) {
		t.Errorf("Test failed: %v != %v\n", actual2, expected2)
	}
}

func TestBin2Dec(t *testing.T) {
	actual1 := Bin2dec([]bool{false, false, true})
	expected1 := 4

	if !(reflect.DeepEqual(actual1, expected1)) {
		t.Errorf("Test failed: %v != %v\n", actual1, expected1)
	}

	actual2 := Bin2dec([]bool{false})
	expected2 := 0

	if !(reflect.DeepEqual(actual2, expected2)) {
		t.Errorf("Test failed: %v != %v\n", actual2, expected2)
	}
}

func TestDivide(t *testing.T) {
	actual1a, actual1b := Divide(testArr2, testArr1)
	expected1a, expected1b := []bool{true, true}, []bool{false}

	if !(reflect.DeepEqual(actual1a, expected1a)) {
		t.Errorf("Test failed: %v != %v\n", actual1a, expected1a)
	}

	if !(reflect.DeepEqual(actual1b, expected1b)) {
		t.Errorf("Test failed: %v != %v\n", actual1b, expected1b)
	}

	actual2a, actual2b := Divide(testArr3, testArr1)
	expected2a, expected2b := []bool{true, true, true, true, true, true}, []bool{true}

	if !(reflect.DeepEqual(actual2a, expected2a)) {
		t.Errorf("Test failed: %v != %v\n", actual2a, expected2a)
	}

	if !(reflect.DeepEqual(actual2b, expected2b)) {
		t.Errorf("Test failed: %v != %v\n", actual2b, expected2b)
	}
}

func TestExp(t *testing.T) {
	actual1 := Exp(testArr4, testArr1)
	expected1 := []bool{false}

	if !(reflect.DeepEqual(actual1, expected1)) {
		t.Errorf("Test failed: %v != %v\n", actual1, expected1)
	}

	actual2 := Exp(testArr1, testArr2)
	expected2 := []bool{false, false, false, false, false, false, true}

	if !(reflect.DeepEqual(actual2, expected2)) {
		t.Errorf("Test failed: %v != %v\n", actual2, expected2)
	}
}

func TestEgcd(t *testing.T) {
	_, _, actual1, _ := Egcd(testArr3, testArr1)
	expected1 := []bool{true}

	if !(reflect.DeepEqual(actual1, expected1)) {
		t.Errorf("Test failed: %v != %v\n", actual1, expected1)
	}

	_, _, actual2, _ := Egcd(testArr1, testArr2)
	expected2 := []bool{false, true}

	if !(reflect.DeepEqual(actual2, expected2)) {
		t.Errorf("Test failed: %v != %v\n", actual2, expected2)
	}
}
