/*
Package bitset implements a fast and simple set that can store small values
(from 0 to 15). Each element is represented by a bit in uint16.
This allows fast set operations.
*/
package bitset

import "fmt"

// Set elements (must be between 0 and 15)
type Element uint8

// Maximum value that can be stored in a set (15).
const MaxValue Element = 15

// Bit set. The zero value of this type is an empty set.
// To check if two sets contain the same elements, just use the == operator.
type BitSet uint16

// Empty set and {0..15}
const (
	EmptySet BitSet = 0
	Universe BitSet = 1<<16 - 1
)

// Set returns a set containing the elements passed in parameter
func Set(er ...Element) BitSet {
	return EmptySet.Add(er...)
}

// Interval returns the set {a ... b}
func Interval(a, b Element) BitSet {
	return (1<<(b-a+1) - 1) << a
}

// Contains returns true if set contains e
func (set BitSet) Contains(e Element) bool {
	return set&(1<<e) != 0
}

// Size returns the size (number of elements) ot a set
func (set BitSet) Size() int {
	s := 0
	for ; set > 0; set >>= 1 {
		s += int(set & 1)
	}
	return s
}

// Min returns the minimal element of a set.
// If the set is empty, returns 0
func (set BitSet) Min() Element {
	if set == EmptySet {
		return 0
	}
	e := Element(0)
	for ; set&1 == 0; set >>= 1 {
		e++
	}
	return e
}

// Max returns the maximum element of a set
// If the set is empty, returns 0
func (set BitSet) Max() Element {
	if set == EmptySet {
		return 0
	}
	e := Element(0)
	for ; set > 0; set >>= 1 {
		e++
	}
	return e - 1
}

func (set BitSet) rangeHelper(r []Element) []Element {
	for e := Element(0); set > 0; e++ {
		if set&1 != 0 {
			r = append(r, e)
		}
		set >>= 1
	}
	return r
}

// Range returns a slice containing the same elements as the set.
// This method is useful for iterationg over a set.
func (set BitSet) Range() []Element {
	return set.rangeHelper(make([]Element, 0, set.Size()))
}

// RangeA works like Range, but returns a slice of the parameter array.
// It does not allocate a new array and is thus faster than Range
func (set BitSet) RangeA(a [16]Element) []Element {
	return set.rangeHelper(a[:0])
}

// String returns a string representation of a set in a form
// [e1 e2 ... en]
func (set BitSet) String() string {
	return fmt.Sprint(set.Range())
}

// Add returns a set containing the same elements as its receiver
// plus the elements in parameter
func (set BitSet) Add(er ...Element) BitSet {
	for _, e := range er {
		set |= 1 << e
	}
	return set
}

// Remove returns a set containing the same element as its receiver
// minus the elements in parameter
func (set BitSet) Remove(er ...Element) BitSet {
	for _, e := range er {
		set &= ^(1 << e)
	}
	return set
}

// Union returns the union of two sets
// One can use set | other instead
func (set BitSet) Union(other BitSet) BitSet {
	return set | other
}

// Intersection returns the intersection of two sets
// One can use set & other instead
func (set BitSet) Intersection(other BitSet) BitSet {
	return set & other
}

// Minus returns the difference of two sets
// One can use set & ^other instead
func (set BitSet) Minus(other BitSet) BitSet {
	return set & ^other
}
