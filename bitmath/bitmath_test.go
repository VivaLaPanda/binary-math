package bitmath

import (
	"math"
	"reflect"
	"testing"
)

// ------- Newer table driven tests. WARNING: These rely on Dec2bin/Bin2dec
// ------- working properly!

var rsaTests = []struct {
	in1 int
}{
	{5},
	{20},
	{100},
}

func TestRsa(t *testing.T) {
	for _, test := range rsaTests {
		rsaD, rsaE, rsaN := RSAKeygen(Dec2bin(10), Dec2bin(20))
		encMsg := Bin2dec(RSAEnc(Dec2bin(test.in1), rsaE, rsaN))
		decMsg := Bin2dec(RSADec(Dec2bin(encMsg), rsaD, rsaN))
		if test.in1 != decMsg {
			t.Errorf("RSA dec(enc(msg)) != msg: enc(%v)=>%v,  dec(enc(%v))=>%v", test.in1, encMsg, test.in1, decMsg)
		}
	}
}

var modInvTests = []struct {
	in_1  int
	in_2  int
	out_1 int
}{
	{5, 17, 7},
	{20, 113, 17},
	{100, 113, 26},
}

func TestModInv(t *testing.T) {
	for _, test := range modInvTests {
		actual := Bin2dec(ModInv(Dec2bin(test.in_1), Dec2bin(test.in_2)))
		if actual != test.out_1 {
			t.Errorf("ModInv(%v, %v) => %v, want %v", test.in_1, test.in_2, actual, test.out_1)
		}
	}
}

var egcdTests = []struct {
	in_1  int
	in_2  int
	out_1 int
	out_2 int
	out_3 int
	out_4 bool
}{ // in2, in2, x, y, gcd, isNegative
	{2, 6, 1, 0, 2, false},
	{20, 113, 17, 3, 1, false},
	{100, 113, 26, 23, 1, false},
	{23, 113, 54, 11, 1, true},
}

func TestEgcd(t *testing.T) {
	for _, test := range egcdTests {
		a_out_1b, a_out_2b, a_out_3b, a_out_4 := Egcd(Dec2bin(test.in_1), Dec2bin(test.in_2))
		a_out_1, a_out_2, a_out_3 := Bin2dec(a_out_1b), Bin2dec(a_out_2b), Bin2dec(a_out_3b)
		if a_out_1 != test.out_1 {
			t.Errorf("Egcd(%v, %v) => %v, %v, %v, %v, want %v, %v, %v, %v", test.in_1,
				test.in_2, a_out_1, a_out_2, a_out_3, a_out_4, test.out_1, test.out_2, test.out_3, test.out_4)
		}
	}
}

// ------- Old tests, not table driven ----------

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

// func TestEgcd(t *testing.T) {
// 	_, _, actual1, _ := Egcd(testArr3, testArr1)
// 	expected1 := []bool{true}
//
// 	if !(reflect.DeepEqual(actual1, expected1)) {
// 		t.Errorf("Test failed: %v != %v\n", actual1, expected1)
// 	}
//
// 	_, _, actual2, _ := Egcd(testArr1, testArr2)
// 	expected2 := []bool{false, true}
//
// 	if !(reflect.DeepEqual(actual2, expected2)) {
// 		t.Errorf("Test failed: %v != %v\n", actual2, expected2)
// 	}
// }

// MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
// MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
// MMMMMMMMMMMMMMMMMMO=OMMMMMMMMMMMMMMMMMMMMNZOMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
// MMMMMMMMMMMMMMMMMM==+=OMMMMMMMMMMMMMMMMMM+?=7MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
// MMMMMMMMMMMMMMMMMO==Z+==8MMMMMMMMMMMMMMMZ=~I==OMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
// MMMMMMMMMMMMMMMMM?Z..7==~OMMMMMMMMMMMMMD?...Z==IMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
// MMMMMMMMMMMMMMMMN$....+=~7O7?========IZO.....====OMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
// MMMMMMMMMMMMMMMMO...7?=================?.....$====8MMMMMMMMMMMMMMMMMMMMMMMMMMMMM
// MMMMMMMMMMMMMMMM8,$~===================+:....Z=====8MMMMMMMMMMMMMMMMMMMMMMMMMMMM
// MMMMMMMMMMMMMMMD7=======================~=?$?Z======ODMMMMMMMMMMMMMMMMMMMMMMMMMM
// MMMMMMMMMMMMMMO====================================~~7ODMMMMMMMMMMMMMMMMMMMMMMMM
// MMMMMMMMMMMMD============================+==========~==O~$MMMMMMMMMMMMMMMMMMMMMM
// MMMMMMMMMMM8==============?============$I=+~=~========$=$M8+8MMMMMMMMMMMMMMMMMMM
// MMMMMMMMMM8========Z===~?I===========~~+~~O==$===========IMMNIDMMMMMMMMMMMMMMMMM
// MMMMMMMMM8~~=====ZI====$~~=======?==~,,Z==$==+============7MMM8ZMMMMMMMMMMMMMMMM
// MMMMMMMMM=.,====Z=O====+O======~Z==~:,,I=I?~=~+=======I~~.,ZMMMMOMMMMMMMMMMMMMMM
// MMMMMMMM$Z=~,:.$~?I=?===Z=======7~~?,,,==ZI=~=O===~==.~I~..=ZMMMMOMMMMMMMMMMMMMM
// MMMMMMMO=======7:=$::$ZI7~,:.==Z==I,,,,+=IO~.=7==,.,:.,$=====MMMMM$MMMMMMMMMMMMM
// MMMMMMMOO~======?,I~~~O~~======?~==,,,,$====~?77========?====OMMMM88MMMMMMMMMMMM
// MMMMMMMZZ======7Z,,:ZZ?I,7$Z$==:==,,,,,==Z,Z===?========Z=====MMMMM+NMMMMMMMMMMM
// MMMMMMMOM===~==+ZO8~..,OO,,,,,Z,$=,,,,Z:??,,O=~$==============8MMMMOOMMMMMMMMMMM
// MMMMMMMMMO~~+==Z8..OOO..+Z,,,,,,,,,,ZOI,:$O$,~?O========~7~~==OMMMMOIMMMMMMMMMMM
// MMMMMMMMMM?~?=~8..8888OO8,,,,,,,,,,,..:ZO$..I8,$=========Z====?MMMM8=MMMMMMMMMMM
// MMMMMMMMMM$Z==?+.OOZ88888,,,,,,,,,,,$8888=....OI=========?=====NMMM8~MMMMMMMMMMM
// MMMMMMMMMD====77.OO7O8OO8,,,,,,,,,,=888888O8...O==============~MMMMO=MMMMMMMMMMM
// MMMMMMMMM?===O,,..8Z$$$OZ,,,,,,,,,,ZZ7888888...O===============MMMM$OMMMMMMMMMMM
// MMMMMMMMO===Z,,,,..?OOO~,,,,,,,,,,,~87$OO$8O..~O=========~+====MMMM=NMMMMMMMMMMM
// MMMMMMMM====::=I=:I,,,,,,,,,,,,,,,,,=88$Z88..,:O==========?===?MMMOOMMMMMMMMMMMM
// MMMMMMMD===O:$~:~7:~,,,,,,,,,,,,,,,,,,,:~,,,,,,Z==========?===OMM8ZMMMMMMMMMMMMM
// MMMMMMMZ===?::II::?,,$.,,,II=,,,,,,,,,:7$~7I:,,7==========+==~8MN8MMMMMMMMMMMMMM
// MMMMMMM?===7,,,,,,,,,O.7IIIIII7I7=7I,,:~?$~$I:$+=========~===?MNMMMMMMMMMMMMMMMM
// MMMMMMM====7,,,,,,,,,IIIIIIIIIIIIIII~,,,~:$~$,I========+=+===OMMMMMMMMMMMMMMMMMM
// MMMMMMM=====7.,,,,,,=III+=~~~~~~~==IO,,,,,,,,=Z========I=$~=+MMMMMMMMMMMMMMMMMMM
// MMMMMMM+=====++,,,,,:I=~~~~~~~~~~~~~7,,,,,,,,IZ========+=Z~=$MMMMMMMMMMMMMMMMMMM
// MMMMMMM$=======Z:,,,,Z~~~~~~~~~~~~~~Z,,,,,,,=~~========~=~==ZMMMMMMMMMMMMMMMMMMM
// MMMMMMM8,.=~=~==~=Z,,,,$=~~~~~~~~~~7,,,,,,,,=$========7=Z~==ZMMMMDDMMMMMMMMMMMMM
// MMMMMMMZ=:~.=,=====I=I$:,,,:I$Z$+,,,,,,,,,7?I7==========+===ZMNDOZ...?.ZMMZMMMMM
// MMMMMMMOO=~========$====Z$ZO$?:,,,,,=ZO$$$+Z=.,======7=Z====$MMMMMO.. ...O7MMMMM
// MMMMMMMNIZ=======I=7==+$$777OZ:,,,,,:$7$OZO$O.,~.~~~~.=~~.~.7MMMMMM$.....~=MMMMM
// MMMMMMMMZMI======$$==I?7$777$8,,,,,=O7$$7O$$========$~I===+==MMMMMMM=~~=~,=$MMMM
// MMMMMMMMM8M7=====Z=+Z++$Z777$OOZZ7I8777$Z7$Z=======?=Z===~~==DMMMMMM+=======OMMM
// MMMMMMMMMMMMO======++++Z7Z77$OZ+++Z777$Z77O~========O=~==~===7MMMMM8========7MMM
// MMMMMMMMMMMMMMO===7+++?$7$O$$$ZO+O$$7$Z$$$I========8====$=====ZMMMM$========IMMM
