package edn

import (
	. "github.com/csm/go-edn/types"
	. "testing"
	"math/big"
)

func parse(s string, t *T) (val Value) {
	val, err := ParseString(s)

	if err != nil {
		t.Errorf("Expected parsing %+v to succeed. %v", s, err)
		t.FailNow()
	}

	return
}

func assertValueEqual(actual, expected Value, t *T) {
	if !expected.Equals(actual) {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}

func TestDoesNotParseEmptyInput(t *T) {
	_, err := ParseString("")

	if err == nil {
		t.Error("Expected parsing empty input to fail")
	}
}

func TestParseList(t *T) {
	val := parse("()", t)

	if !new(List).Equals(val) {
		t.Errorf("Expected parsing \"()\" to return an empty list, got %+v", val)
	}

	str := `(() "abc" [] "def")`
	val = parse(str, t)

	l := new(List)
	l.Insert(new(List))
	l.Insert(String("abc"))
	l.Insert(Vector{})
	l.Insert(String("def"))

	if !l.Equals(val) {
		t.Errorf("Expected %v, got %v", l, val)
	}
}

func TestParseVector(t *T) {
	assertValueEqual(parse(`[]`, t), Vector{}, t)
	assertValueEqual(parse(`[[]]`, t), Vector{Vector{}}, t)
	assertValueEqual(parse(`[[("abc")] "def"]`, t),
		Vector{
			Vector{
				new(List).Insert(String("abc")),
			},
			String("def"),
		}, t,
	)
}

func TestParseMap(t *T) {
	_, err := ParseString("{:x}")

	if err == nil || err.Error() != "Error: Map literal must contain an even number of forms" {
		t.Error("Expected parsing map with single item to fail")
	}

	assertValueEqual(parse(`{}`, t), Map{}, t)
	assertValueEqual(parse(`{:a :b}`, t), Map{Keyword("a"): Keyword("b")}, t)

}

func TestParseCharacter(t *T) {
	assertValueEqual(parse(`\r`, t), Character('r'), t)
	assertValueEqual(parse(`\\`, t), Character('\\'), t)
	assertValueEqual(parse(`\~`, t), Character('~'), t)
	assertValueEqual(parse(`\!`, t), Character('!'), t)
	assertValueEqual(parse(`\[`, t), Character('['), t)
	assertValueEqual(parse(`\5`, t), Character('5'), t)
	assertValueEqual(parse(`\space`, t), Character(' '), t)
	assertValueEqual(parse(`\return`, t), Character('\r'), t)
	assertValueEqual(parse(`\newline`, t), Character('\n'), t)
	assertValueEqual(parse(`\tab`, t), Character('\t'), t)
}

func TestParseString(t *T) {
	val := parse(`""`, t)
	if val == nil || !val.Equals(String("")) {
		t.Errorf("Expected \"\", got %v", val)
	}

	val = parse(`"abc"`, t)
	if val == nil || !val.Equals(String("abc")) {
		t.Errorf("Expected \"abc\", got %v", val)
	}
}

func TestParse(t *T) {
	expected := new(List).Insert(
		String("abc"),
		new(List).Insert(String("spaced")),
		Vector{
			String("vec"),
			new(List).Insert(
				String("an"),
				String("inner"),
				String("list"),
			),
		},
		new(List),
		Vector{},
		new(List),
		String(""),
		Vector{Character('x'), Character('\n')},
		Map{},
		Set{}.Insert(String("set")),
		Value(nil),
		Map{String("key"): String("value")},
		Map{
			Keyword("int"): Int(31337),
			Keyword("float"): Float(2.718),
			Keyword("bigint"): (*BigInt)(big.NewInt(33333)),
			Keyword("rational"): (*Rational)(big.NewRat(123, 456)),
		},
		Map{
			String("key"):   String("value"),
			Keyword("key2"): new(List).Insert(Keyword("value")),
		},
		Keyword("key"),
	)

	actual := parse(`
	("abc" ( "spaced" )
	    ["vec"( "an"	"inner""list",)]
		(),[]()""[\x \newline]
		{}#{"set"} nil
		{"key" "value"}
		{:int 31337 :float 2.718 :bigint 33333N :rational 123/456 }
		{"key" "value" :key2 (:value)}
		:key
	)
	`, t)

	assertValueEqual(actual, expected, t)
}

func TestParseInt(t *T) {
	assertValueEqual(parse("0", t), Int(0), t)
	assertValueEqual(parse("1234", t), Int(1234), t)
}

func TestParseBigInt(t *T) {
	assertValueEqual(parse("31337N", t), (*BigInt)(big.NewInt(31337)), t)
}

func TestParseFloat(t *T) {
	assertValueEqual(parse("1234M", t), Float(1234.0), t)
	assertValueEqual(parse("3.141", t), Float(3.141), t)
}

func TestParseRational(t *T) {
	assertValueEqual(parse("3/4", t), (*Rational)(big.NewRat(3, 4)), t)
}

func TestParseSymbol(t *T) {
	assertValueEqual(parse("a-symbol", t), Symbol("a-symbol"), t)
}

func TestTaggedValue(t *T) {
	assertValueEqual(parse("#tagged \"value\"", t), TaggedValue{"tagged", String("value")}, t)
}