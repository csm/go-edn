package types

import (
	. "testing"
	"fmt"
)

func assertEqual(expect, actual interface{}, t *T) {
	if expect != actual {
		t.Errorf("Expecting %+v, received %+v", expect, actual)
	}
}

func TestVectorString(t *T) {
	vec := make(Vector, 0)
	assertEqual("[]", vec.String(), t)

	vec = append(vec, Int(1))
	assertEqual("[1]", vec.String(), t)

	vec = append(vec, make(Vector, 0), String("abc"))
	assertEqual(`[1 [] "abc"]`, vec.String(), t)
}

func TestListString(t *T) {
	list := new(List)
	ll := list.raw()

	assertEqual("()", list.String(), t)

	ll.PushBack(Int(1))
	assertEqual("(1)", list.String(), t)

	ll.PushBack(make(Vector, 0))
	ll.PushBack(String("abc"))
	assertEqual("(1 [] \"abc\")", list.String(), t)
}

func TestMapString(t *T) {
	_map := make(Map)
	assertEqual("{}", _map.String(), t)

	_map[String("test")] = Vector{String("value1"), String("value2")}
	assertEqual(`{"test" ["value1" "value2"]}`, _map.String(), t)
}

func TestSetString(t *T) {
	set := make(Set)
	assertEqual("#{}", set.String(), t)

	set.Insert(String("abc"))
	assertEqual(`#{"abc"}`, set.String(), t)

	// Unstable test because sets/maps are unordered and enumerated in a
	// random order:
	//
	// set.Insert(Int(123))
	// assertEqual(`#{"abc" 123}`, set.String(), t)

	// TODO: This causes a runtime panic. Mutable types (maps, vectors, etc) can't
	//       be map keys in Go. Since the sets are backed by a Map, it blows up :(
	//
	//       Look into https://code.google.com/p/gohash/
	//
	// set.Insert(Vector{Int(1), Int(2), Map{String("foo"): Int(17)}})
}

func assertValueEqual(v1, v2 Value, t *T) {
	if !v1.Equals(v2) {
		t.Errorf("Expecting %+v, got %+v", v1, v2)
	}
}

func assertValueNotEqual(v1, v2 Value, t*T) {
	if v1.Equals(v2) {
		t.Errorf("expecting %+v not equal to %+v", v1, v2)
	}
}

func TestStrings(t *T) {
	assertValueEqual(String("foobar"), String(fmt.Sprint("foo", "bar")), t)
	assertValueNotEqual(String("foo"), Keyword("foo"), t)
}