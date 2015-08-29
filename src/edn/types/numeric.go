package types

import (
	"fmt"
	"math/big"
	"strings"
	"strconv"
)

// TODO: type Decimal (representing EDN fixed point numbers)

// Int represents EDN integers: 1, -5, 1000.
type Int int64

// Float represents EDN floats: 1.234, -1.2925e+9
type Float float64

// Rational represents EDN rationals: 1/100, -7/5
type Rational big.Rat

// BigInt represents EDN arbitrary precision integers.
type BigInt big.Int

type Bool bool

// Equals compares the Float to another Value for equality.
func (this Float) Equals(other Value) bool {
	return this == other
}

// String returns the EDN string representation of the Float.
func (this Float) String() string {
	return fmt.Sprintf("%f", float64(this))
}

// Equals compares the Int to another Value for equality.
func (this Int) Equals(other Value) bool {
	return this == other
}

// String returns the EDN string representation of the Int.
func (this Int) String() string {
	return fmt.Sprintf("%d", int64(this))
}

func (this Bool) Equals(other Value) bool {
	return this == other
}

func (this Bool) String() string {
	if this {
		return "true"
	} else {
		return "false"
	}
}

func (this *BigInt) Equals(other Value) bool {
	switch o := other.(type) {
	case *BigInt: return (*big.Int)(this).Cmp((*big.Int)(o)) == 0
	default: return false
	}
}

func (this *BigInt) String() string {
	return fmt.Sprint((*big.Int)(this).String(), "N")
}

func (this *Rational) Equals(other Value) bool {
	switch o := other.(type) {
	case *Rational: return (*big.Rat)(this).Cmp((*big.Rat)(o)) == 0
	default: return false
	}
}

func (this *Rational) String() string {
	return (*big.Rat)(this).String()
}

func ParseBigInt(s string) *BigInt {
	var i = big.NewInt(0)
	_, _ = i.SetString(strings.Trim(s, "N"), 10)
	return (*BigInt)(i)
}

func ParseInt(s string) Int {
	var res, _ = strconv.ParseInt(s, 10, 64)
	return Int(res)
}

func ParseFloat(s string) Float {
	var res, _ = strconv.ParseFloat(strings.Trim(s, "M"), 64)
	return Float(res)
}

func ParseRational(s string) *Rational {
	p := strings.Split(s, "/")
	n, _ := strconv.ParseInt(p[0], 10, 64)
	d, _ := strconv.ParseInt(p[1], 10, 64)
	return (*Rational)(big.NewRat(n, d))
}