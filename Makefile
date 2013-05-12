.PHONY: all setup test clean parser lexer get-deps

all: test
	go build -v ...

setup: parser lexer
	@:

lexer:
	nex -s < lexer.nn > lexer.nn.go
	go fmt lexer.nn.go

parser:
	rm -f parser.y.go
	go tool yacc -o parser.y.go parser.y

get-deps:
	go get -d -v -u ... github.com/blynn/nex
	go install github.com/blynn/nex

clean:
	rm -rf *.output *.nn.go *.y.go

test: clean setup
	go test -v
