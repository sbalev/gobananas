package bitset

import (
	"fmt"
	"testing"
)

func TestSize(t *testing.T) {
	sets := []BitSet{EmptySet, Universe, Interval(5, 9)}
	sizes := []int{0, int(MaxValue) + 1, 5}
	for i := 0; i < 3; i++ {
		if sets[i].Size() != sizes[i] {
			t.Errorf("%v.Size() = %d, want %d",
				sets[i], sets[i].Size(), sizes[i])
		}
	}
}

func TestContains(t *testing.T) {
	s := Interval(7, 11)
	var e Element
	for e = 0; e <= MaxValue; e++ {
		if (7 <= e && e <= 11) != s.Contains(e) {
			t.Errorf("%v.Contains(%d) = %t, want %t",
				s, e, s.Contains(e), 7 <= e && e <= 11)
		}
	}
}

func TestAddRemove(t *testing.T) {
	s := EmptySet.Add(3, 4, 5, 6, 7, 8)
	if s != Interval(3, 8) {
		t.Errorf("s = %v, wanted %v", s, Interval(3, 8))
	}
	s = s.Remove(3, 8)
	if s != Interval(4, 7) {
		t.Errorf("s = %v, wanted %v", s, Interval(4, 7))
	}
}

func TestRange(t *testing.T) {
	var s BitSet
	u := Interval(6, 9)
	for _, e := range u.Range() {
		s = s.Add(e)
	}
	if s != u {
		t.Errorf("got %v, wanted %v", s, u)
	}
}

func TestIntersectionUnionMinus(t *testing.T) {
	var odd, even BitSet
	var e Element
	for e = 0; e <= MaxValue; e++ {
		if e&1 == 1 {
			odd = odd.Add(e)
		} else {
			even = even.Add(e)
		}
	}
	s := odd.Intersection(even)
	if s != EmptySet {
		t.Errorf("got %v, wanted %v", s, EmptySet)
	}
	s = odd.Union(even)
	if s != Universe {
		t.Errorf("got %v, wanted %v", s, Universe)
	}
	s = Universe.Minus(odd)
	if s != even {
		t.Errorf("got %v, wanted %v", s, even)
	}
}

func TestMinMax(t *testing.T) {
	if EmptySet.Min() != 0 {
		t.Errorf("got %d, wanted 0", EmptySet.Min())
	}
	if EmptySet.Max() != 0 {
		t.Errorf("got %d, wanted 0", EmptySet.Max())
	}
	s := EmptySet.Add(4, 3, 7, 6)
	if s.Min() != 3 {
		t.Errorf("got %d, wanted 3", s.Min())
	}
	if s.Max() != 7 {
		t.Errorf("got %d, wanted 7", s.Max())
	}
}

func ExampleBitSet() {
	var evens BitSet
	var e Element
	for e = 0; e < MaxValue; e += 2 {
		evens = evens.Add(e)
	}
	fmt.Println(evens)
	odds := Universe.Minus(evens)
	fmt.Println(odds)
	primes := odds.Add(2).Remove(1, 9, 15)
	fmt.Println(primes, primes.Size())
	evenPrimes := evens.Intersection(primes)
	for _, p := range evenPrimes.Range() {
		fmt.Println(p, "is even prime")
	}
	fmt.Println(Interval(5, 9))
	// Output:
	// [0 2 4 6 8 10 12 14]
	// [1 3 5 7 9 11 13 15]
	// [2 3 5 7 11 13] 6
	// 2 is even prime
	// [5 6 7 8 9]
}

func BenchmarkRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Universe.Range()
	}
}

func BenchmarkRangeA(b *testing.B) {
	var a [16]Element
	for i := 0; i < b.N; i++ {
		_ = Universe.RangeA(a)
	}
}
