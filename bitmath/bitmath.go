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
		return []bool{}
	}
	m := n / 2
	A := dec2bin(m)
	fbit := n % 2
	return append([]bool{fbit != 0}, A...)
}

func bin2dec(n []bool) int {
	if len(n) <= 3 {
		return binDecMap(n)
	}

	temp1, temp2 := divide(n, []bool{false, true, false, true})
	return bin2dec(trim(temp1)) + binDecMap(trim(temp2))
}

func divide(X []bool, Y []bool) (q []bool, r []bool) {
	if zero(X) {
		return emptySlice, emptySlice
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

func sliceEq(a []bool, b []bool) bool {
	for i, aElem := range a {
		if aElem != b[i] {
			return false
		}
	}

	return true
}

func binDecMap(v []bool) int {
	switch {
	case sliceEq(v, []bool{}):
		return '0'
	case sliceEq(v, []bool{false}):
		return '0'
	case sliceEq(v, []bool{true}):
		return '1'
	case sliceEq(v, []bool{false, true}):
		return '2'
	case sliceEq(v, []bool{true, true}):
		return '3'
	case sliceEq(v, []bool{false, false, true}):
		return '4'
	case sliceEq(v, []bool{true, false, true}):
		return '5'
	case sliceEq(v, []bool{false, true, true}):
		return '6'
	case sliceEq(v, []bool{true, true, true}):
		return '7'
	case sliceEq(v, []bool{false, false, false, true}):
		return '8'
	case sliceEq(v, []bool{true, false, false, true}):
		return '9'
	}

	panic("FATAL ERROR: Map reached unreachable code.")
}

// TODO: Implement
func sub(A []bool, B []bool) (diff []bool) {
	return []bool{false}
}

func compare(A []bool, B []bool) int {
	A1 := reverse(trim(A))
	A2 := reverse(trim(B))
	if len(A1) > len(A2) {
		return 1
	} else if len(A1) < len(A2) {
		return 2
	}

	for i, _ := range A1 {
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
	// Make a copy of the slices passed in
	A1 := make([]bool, len(A))
	copy(A1, A)
	B1 := make([]bool, len(B))
	copy(B1, B)

	// Padding to equal length
	n := len(A1)
	m := len(B1)
	if n < m {
		dif := len(B1) - len(A1)
		for i := 0; i < dif; i++ {
			A1 = append(A1, false)
		}
	} else {
		dif := len(A1) - len(B1)
		for i := 0; i < dif; i++ {
			B1 = append(B1, false)
		}
	}

	var N int
	if n > m {
		N = n
	} else {
		N = m
	}

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
