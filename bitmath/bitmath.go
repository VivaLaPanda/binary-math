package bitmath

// Useful slices
var emptySlice = []bool{}
var zeroSlice = []bool{false}
var oneSlice = []bool{true}

func Mult(X int, Y int) int {
	X1 := dec2bin(X)
	X2 := dec2bin(Y)
	return bin2dec(mult(X1, X2))
}

func Add(A int, B int) int {
	return bin2dec(add(dec2bin(A), dec2bin(B)))
}

func Divide(X int, Y int) (int, int) {
	q, r := divide(dec2bin(X), dec2bin(Y))
	return bin2dec(q), bin2dec(r)
}

func dec2bin(n int) []bool {
	if n == 0 {
		return []bool{false}
	}
	m := n / 2
	A := dec2bin(m)
	fbit := n % 2
	return append([]bool{fbit != 0}, A...)
}

func bin2dec(A []bool) int {
	if len(A) == 0 {
		return 0
	}

	var val int
	if A[0] {
		val = 1
	} else {
		val = 0
	}

	pow := 2
	for i := 1; i < len(A); i++ {
		var currVal int
		if A[i] {
			currVal = 1
		} else {
			currVal = 0
		}

		val = val + pow*currVal
		pow = pow * 2
	}

	return val
}

// DEPRECATED
// func bin2dec1(n []bool) int {
// 	if len(n) <= 3 {
// 		return binDecMap(trim(n))
// 	}
//
// 	temp1, temp2 := divide(n, []bool{false, true, false, true})
// 	return bin2dec1(trim(temp1)) + binDecMap(trim(temp2))
// }

func divide(X []bool, Y []bool) (q []bool, r []bool) {
	if zero(X) {
		return zeroSlice, zeroSlice
	}

	q, r = divide(div2(X), Y)
	q = add(q, q)
	r = add(r, r)
	if !even(X) {
		r = add(r, oneSlice)
	}
	if !(compare(r, Y) == 2) {
		r = sub(r, Y)
		q = add(q, oneSlice)
	}
	return q, r
}

// Faster than deepeq for production code
func sliceEq(a []bool, b []bool) bool {
	if len(a) != len(b) {
		return false
	}

	for i, aElem := range a {
		if aElem != b[i] {
			return false
		}
	}

	return true
}

// DEPRECATED
// func binDecMap(v []bool) int {
// 	switch {
// 	case sliceEq(v, []bool{}):
// 		return '0'
// 	case sliceEq(v, []bool{false}):
// 		return '0'
// 	case sliceEq(v, []bool{true}):
// 		return '1'
// 	case sliceEq(v, []bool{false, true}):
// 		return '2'
// 	case sliceEq(v, []bool{true, true}):
// 		return '3'
// 	case sliceEq(v, []bool{false, false, true}):
// 		return '4'
// 	case sliceEq(v, []bool{true, false, true}):
// 		return '5'
// 	case sliceEq(v, []bool{false, true, true}):
// 		return '6'
// 	case sliceEq(v, []bool{true, true, true}):
// 		return '7'
// 	case sliceEq(v, []bool{false, false, false, true}):
// 		return '8'
// 	case sliceEq(v, []bool{true, false, false, true}):
// 		return '9'
// 	}
//
// 	panic("FATAL ERROR: Map reached unreachable code.")
// }

func equalLengthPad(ARef []bool, BRef []bool) (A []bool, B []bool, N int) {
	// Make a copy of the slices passed in
	A = make([]bool, len(ARef))
	copy(A, ARef)
	B = make([]bool, len(BRef))
	copy(B, BRef)

	// Padding to equal length
	n := len(A)
	m := len(B)
	if n < m {
		dif := len(B) - len(A)
		for i := 0; i < dif; i++ {
			A = append(A, false)
		}
	} else {
		dif := len(A) - len(B)
		for i := 0; i < dif; i++ {
			B = append(B, false)
		}
	}

	if n > m {
		N = n
	} else {
		N = m
	}

	return A, B, N
}

func twosCompliment(ARef []bool) (A []bool) {
	// Make a copy of the slices passed in
	A = make([]bool, len(ARef))
	copy(A, ARef)

	// 2's compliment
	// Inverting B
	for i, elem := range A {
		A[i] = !elem
	}
	// Add 1
	return add(A, oneSlice)
}

func sub(ARef []bool, BRef []bool) (diff []bool) {
	A, B, n := equalLengthPad(ARef, BRef)

	// Add one extra 0 in front of each to make sure 2compliment has sign bit
	A = append(A, false)
	B = append(B, false)

	signCheck := compare(A, B)
	if signCheck == 0 {
		return zeroSlice
	}

	B = twosCompliment(B)
	diff = add(A, B)

	// Get rid of any overflow if it happened
	diff = diff[:n+1]

	if diff[n] == true {
		return trim(twosCompliment(diff))
	} else {
		return trim(diff)
	}
}

func compare(A []bool, B []bool) int {
	A1 := reverse(trim(A))
	A2 := reverse(trim(B))
	if len(A1) > len(A2) {
		return 1
	} else if len(A1) < len(A2) {
		return 2
	}

	for i := range A1 {
		if A1[i] && !A2[i] {
			return 1
		} else if !A1[i] && A2[i] {
			return 2
		}
	}

	return 0
}

func reverse(A []bool) []bool {
	B := make([]bool, len(A))
	copy(B, A)
	last := len(B) - 1
	for i := 0; i < len(B)/2; i++ {
		B[i], B[last-i] = B[last-i], B[i]
	}

	return B
}

func trim(A []bool) []bool {
	if len(A) == 0 {
		return A
	}
	A1 := reverse(A)
	for !(len(A1) == 0) && (A1[0] == false) {
		A1 = A1[1:]
	}

	return reverse(A1)
}

func shift(A []bool, n int) []bool {
	if n == 0 {
		return A
	}

	return append(zeroSlice, shift(A, n-1)...)
}

// Modified egyptian algorithm
func mult(X []bool, Y []bool) []bool {
	if zero(Y) {
		return zeroSlice
	} else if zero(X) {
		return zeroSlice
	}

	Z := mult(X, div2(Y))

	if even(Y) {
		return add(Z, Z)
	}

	return add(X, add(Z, Z))
}

func zero(X []bool) bool {
	if len(X) == 0 {
		return true
	}

	for _, elem := range X {
		if elem == true {
			return false
		}
	}

	return true
}

func div2(Y []bool) []bool {
	if len(Y) == 0 {
		return Y
	}

	return Y[1:]
}

func even(X []bool) bool {
	if (len(X) == 0) || (X[0] == false) {
		return true
	}

	return false
}

func add(A []bool, B []bool) []bool {
	A1, B1, N := equalLengthPad(A, B)

	C := []bool{}
	carry := false
	for i := 0; i < N; i++ {
		C = append(C, excOr(A1[i], B1[i], carry))
		carry = nextcarry(carry, A1[i], B1[i])
	}
	if carry == true {
		C = append(C, carry)
	}

	return C
}

func excOr(a bool, b bool, c bool) bool {
	return (a != (b != c))
}

func nextcarry(a bool, b bool, c bool) bool {
	if (a && b) || (b && c) || (c && a) {
		return true
	}

	return false
}
