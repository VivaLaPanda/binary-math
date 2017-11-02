package bitmath

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

// Useful slices
var emptySlice = []bool{}
var zeroSlice = []bool{false}
var oneSlice = []bool{true}

// Problems

//
// Actual core computations
//

// --- Integer Operations --

func IntMult(X int, Y int) int {
	X1 := Dec2bin(X)
	X2 := Dec2bin(Y)
	return Bin2dec(Mult(X1, X2))
}

func IntAdd(A int, B int) int {
	return Bin2dec(Add(Dec2bin(A), Dec2bin(B)))
}

func IntSub(A int, B int) (diff int) {
	binDiff, isNegative := Sub(Dec2bin(A), Dec2bin(B))
	diff = Bin2dec(binDiff)
	if isNegative {
		diff = -1 * diff
	}

	return diff
}

func IntExp(A int, B int) int {
	return Bin2dec(Exp(Dec2bin(A), Dec2bin(B)))
}

func IntDivide(X int, Y int) (int, int) {
	q, r := Divide(Dec2bin(X), Dec2bin(Y))
	return Bin2dec(q), Bin2dec(r)
}

func Dec2bin(n int) []bool {
	if n == 0 {
		return []bool{false}
	}

	m := n / 2
	A := Dec2bin(m)
	fbit := n % 2
	return append([]bool{fbit != 0}, A...)
}

func Bin2dec(A []bool) int {
	if len(A) == 0 {
		return 0
	} else if len(A) > 64 {
		fmt.Println("FATAL ERR: Input is greater than 64 bit, user BigBin2dec")
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

func BigDec2bin(n big.Int) []bool {
	if n.Cmp(new(big.Int).SetInt64(0)) == 0 {
		return []bool{false}
	}
	m := new(big.Int).SetInt64(2)
	m.Div(&n, m)
	A := BigDec2bin(*m)
	fbit := m
	m.Mod(&n, new(big.Int).SetInt64(2))
	boolFBit := (fbit.Cmp(new(big.Int).SetInt64(0)) != 0)
	return append([]bool{boolFBit}, A...)
}

func BigBin2dec(A []bool) big.Int {
	if len(A) == 0 {
		return *new(big.Int).SetInt64(0)
	}

	var val big.Int
	if A[0] {
		val = *new(big.Int).SetInt64(1)
	} else {
		val = *new(big.Int).SetInt64(0)
	}

	pow := new(big.Int).SetInt64(2)
	for i := 1; i < len(A); i++ {
		var currVal big.Int
		if A[i] {
			currVal = *new(big.Int).SetInt64(1)
		} else {
			currVal = *new(big.Int).SetInt64(0)
		}

		temp := new(big.Int).SetInt64(0)
		val.Add(&val, temp.Mul(pow, &currVal))
		pow.Mul(pow, new(big.Int).SetInt64(2))
	}

	return val
}

// --- Binary Operations --

// --- Binary Operations --
func RSAKeygen(nBits []bool, confidence []bool) (rsaD, rsaE, rsaN []bool) {
	var rsaP []bool
	var rsaQ []bool

	rsaP = NBitPrime(div2(nBits), confidence)
	rsaQ = NBitPrime(div2(nBits), confidence)
	rsaN = Mult(rsaP, rsaQ)
	temp1, _ := Sub(rsaP, oneSlice)
	temp2, _ := Sub(rsaQ, oneSlice)
	thOfN := Mult(temp1, temp2)
	rsaE = []bool{true, true} // 3
	// Make sure E is relatively prime with theOfN
	_, _, d, _ := Egcd(thOfN, rsaE)
	if compare(d, oneSlice) != 0 {
		rsaE = []bool{true, true, true} // 7
		_, _, d, _ := Egcd(thOfN, rsaE)
		if compare(d, oneSlice) != 0 {
			lessBits, _ := Sub(nBits, []bool{false, true}) // Simplest way to ensure e < thOfN
			rsaE = NBitPrime(lessBits, confidence)
			_, _, d, _ := Egcd(thOfN, rsaE)

			for compare(d, oneSlice) != 0 {
				rsaE = NBitPrime(lessBits, confidence)
				_, _, d, _ = Egcd(thOfN, rsaE)
			}
		}
	}

	rsaD = ModInv(rsaE, thOfN)

	return rsaD, rsaE, rsaN
}

func RSAEnc(msg, rsaE, rsaN []bool) (encMsg []bool) {
	_, encMsg = Divide(Exp(msg, rsaE), rsaN)
	return encMsg
}

func RSADec(encMsg, rsaD, rsaN []bool) (msg []bool) {
	_, msg = Divide(Exp(encMsg, rsaD), rsaN)
	return msg
}

func NBitPrime(n []bool, confidence []bool) (prime []bool) {
	for isPrime := false; !isPrime; {
		// Seed the random number generator
		s1 := rand.NewSource(time.Now().UnixNano())
		seedRand := rand.New(s1)

		prime = make([]bool, 0)
		for i := zeroSlice; compare(i, n) == 2; i = Add(i, oneSlice) {
			randbit := (seedRand.Intn(2) == 1)
			prime = append(prime, randbit)
		}

		prime = prime[2:]
		prime = append(append(oneSlice, prime...), oneSlice...)

		isPrime = PrimalityThree(prime, confidence)
	}

	return prime
}

func ModInv(a []bool, n []bool) (inv []bool) {
	x, _, d, isNegative := Egcd(a, n)
	if compare(d, oneSlice) != 0 {
		return zeroSlice
	}
	quot, inv := Divide(x, n) // mod
	if !isNegative {
		return inv
	}

	inv, _ = Sub(Mult(Add(quot, oneSlice), n), x)
	return inv
}

func PrimalityThree(n []bool, confidence []bool) (isPrime bool) {
	_, divisible2 := Divide(n, Dec2bin(2))
	_, divisible3 := Divide(n, Dec2bin(3))
	_, divisible5 := Divide(n, Dec2bin(5))
	_, divisible7 := Divide(n, Dec2bin(7))

	if zero(divisible2) || zero(divisible3) || zero(divisible5) || zero(divisible7) {
		return false
	}

	return primalityTwo(n, confidence)
}

func primalityTwo(n []bool, confidence []bool) (isPrime bool) {
	for i := zeroSlice; compare(i, confidence) == 2; i = Add(i, oneSlice) {
		if !primality(n) {
			return false
		}
	}

	return true
}

func primality(n []bool) (isPrime bool) {
	a := randBin(oneSlice, n)
	nMinusOne, _ := Sub(n, oneSlice)
	_, r := Divide(Exp(a, nMinusOne), n)

	return compare(r, oneSlice) == 0
}

func randBin(min []bool, max []bool) (randNum []bool) {
	randNum = make([]bool, len(max))
	// Seed the random number generator
	s1 := rand.NewSource(time.Now().UnixNano())
	seedRand := rand.New(s1)

	for i := range max {
		// Set all bits in rand to random bools (gen number 0/1 and then cast to bool)
		randbit := (seedRand.Intn(2) == 1)
		randNum[i] = randbit
	}

	// Ensure the number is smaller than max
	_, randNum = Divide(randNum, max)

	// If the number is smaller than min (really unlikely) just recalc
	if compare(randNum, min) == 2 {
		return randBin(min, max)
	} else {
		return randNum
	}
}

// Computes x^y
func Exp(x []bool, y []bool) []bool {
	if sliceEq(y, zeroSlice) || sliceEq(y, emptySlice) {
		return oneSlice
	}

	z := Exp(x, div2(y))
	if even(y) {
		return Mult(z, z)
	}

	return Mult(x, Mult(z, z))
}

// Takes a and b and finds integers s.t. d = gcd(a,b) and ax + by = d
// Currently cannot distriguish negative answers properly.
func Egcd(a []bool, b []bool) (x []bool, y []bool, d []bool, isNegative bool) {
	if zero(b) {
		return oneSlice, zeroSlice, a, false
	}

	q, r := Divide(a, b)
	x, y, d, isNegative = Egcd(b, r) // in this case isNegative refers to y
	x = Add(x, Mult(q, y))
	isNegative = !isNegative
	return y, x, d, isNegative
}

// DEPRECATED
// func Bin2dec1(n []bool) int {
// 	if len(n) <= 3 {
// 		return binDecMap(trim(n))
// 	}
//
// 	temp1, temp2 := divide(n, []bool{false, true, false, true})
// 	return Bin2dec1(trim(temp1)) + binDecMap(trim(temp2))
// }

func Divide(X []bool, Y []bool) (q []bool, r []bool) {
	if zero(X) {
		return zeroSlice, zeroSlice
	}

	q, r = Divide(div2(X), Y)
	q = Add(q, q)
	r = Add(r, r)
	if !even(X) {
		r = Add(r, oneSlice)
	}
	if !(compare(r, Y) == 2) {
		r, _ = Sub(r, Y)
		q = Add(q, oneSlice)
	}
	return q, r
}

func Sub(ARef []bool, BRef []bool) (diff []bool, isNegative bool) {
	A, B, n := equalLengthPad(ARef, BRef)

	// Add one extra 0 in front of each to make sure 2compliment has sign bit
	A = append(A, false)
	B = append(B, false)

	signCheck := compare(A, B)
	if signCheck == 0 {
		return zeroSlice, false
	}

	B = twosCompliment(B)
	diff = Add(A, B)

	// Get rid of any overflow if it happened
	diff = diff[:n+1]

	if diff[n] == true {
		return trim(twosCompliment(diff)), true
	} else {
		return trim(diff), false
	}
}

// Modified egyptian algorithm
func Mult(X []bool, Y []bool) []bool {
	if zero(Y) {
		return zeroSlice
	} else if zero(X) {
		return zeroSlice
	}

	Z := Mult(X, div2(Y))

	if even(Y) {
		return Add(Z, Z)
	}

	return Add(X, Add(Z, Z))
}

func Add(A []bool, B []bool) []bool {
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

func excOr(a bool, b bool, c bool) bool {
	return (a != (b != c))
}

func nextcarry(a bool, b bool, c bool) bool {
	if (a && b) || (b && c) || (c && a) {
		return true
	}

	return false
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
	return Add(A, oneSlice)
}
