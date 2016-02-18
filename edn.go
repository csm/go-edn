//go:generate go get github.com/blynn/nex
//go:generate go install github.com/blynn/nex
//go:generate go run $GOPATH/src/github.com/blynn/nex/main.go $GOPATH/src/github.com/blynn/nex/nex.go -o lexer.nn.go lexer.nn
//go:generate go fmt lexer.nn.go
//go:generate sed -i~ s:Lexer:lexer:g lexer.nn.go
//go:generate sed -i~ s:Newlexer:newLexer:g lexer.nn.go
//go:generate go tool yacc -o parser.y.go parser.y
//go:generate sed -i~ -f fixparser.sed parser.y.go

package edn

import (
	"errors"
	"fmt"
	"github.com/csm/go-edn/types"
	"io"
	"strings"
)

// ParseString is like ParseReader except it takes a string.
// See ParseReader for more details.
func ParseString(string string) (val types.Value, err error) {
	val, err = ParseReader(strings.NewReader(string))
	return
}

// ParseReader parses EDN from an io.Reader.
//
// Data is returned as a Value in the first return value.
// The second return value is nil on successful parses, and
// an error on unsuccessful parses (e.g. syntax error).
func ParseReader(reader io.Reader) (val types.Value, err error) {
	defer func() {
		// Nex's parser calls panic() on a lexing error
		if r := recover(); r != nil {
			if err == nil {
				err = errors.New(fmt.Sprintf("Error: %v", r))
			}
		}
	}()

	lexer := newLexer(reader)
	result := new(yySymType)
	if yyParse(lexer, result) == 0 {
		//fmt.Printf("result: v:%T:%+v\n", result.v, result.v)
		val = result.v
	} else {
		err = errors.New("Error: could not parse provided EDN")
	}

	return
}

// DumpString accepts any EDN value and will return the EDN string
// representation.
func DumpString(value types.Value) string {
	return value.String()
}
