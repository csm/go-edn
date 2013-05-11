package edn

import . "testing"
import ll "container/list"

func assertEqual(expect, actual interface{}, t *T) {
	if expect != actual {
		t.Errorf("Expecting %+v, received %+v", expect, actual)
	}
}

func TestVectorString(t *T) {
	vec := make(Vector, 0)
	str := vec.String()

	if str != "[]" {
		t.Fail()
	}

	vec = append(vec, Int(1))
	assertEqual("[1]", vec.String(), t)

	vec = append(vec, make(Vector, 0), String("abc"))
	assertEqual(`[1 [] "abc"]`, vec.String(), t)
}

func TestListString(t *T) {
	list := new(List)
	ll := (*ll.List)(list)

	assertEqual("()", list.String(), t)

	ll.PushBack(Int(1))
	assertEqual("(1)", list.String(), t)

	ll.PushBack(make(Vector, 0))
	ll.PushBack(String("abc"))
	assertEqual("(1 [] \"abc\")", list.String(), t)
}

func TestEmptyMapString(t *T) {
	_map := new(Map)
	str := _map.String()

	if str != "{}" {
		t.Fail()
	}
}

func TestEmptySetString(t *T) {
	set := new(Set)
	str := set.String()

	if str != "#{}" {
		t.Fail()
	}
}