.PHONY: all prepare test clean parser lexer get-deps

all: prepare test
	go build -v edn

prepare: parser lexer
	@:

lexer:
	./bin/nex < src/edn/lexer.nn > src/edn/lexer.nn.go
	go fmt src/edn/lexer.nn.go

	# Make "Lexer" private by renaming it... :S
	sed  -i '' 's:Lexer:lexer:g' src/edn/lexer.nn.go
	sed  -i '' 's:Newlexer:newLexer:g' src/edn/lexer.nn.go

parser:
	rm -f src/edn/parser.y.go
	go tool yacc -o src/edn/parser.y.go src/edn/parser.y

	# Make "yyParse" accept a reference to a result
	sed -i '' 's:yyParse(yylex yyLexer) int:yyParse(yylex yyLexer, result *yySymType) int:' src/edn/parser.y.go

get-deps:
	go get -d -v -u github.com/blynn/nex
	go install github.com/blynn/nex 

clean:
	rm -rf *.output *.nn.go *.y.go

test:
	go test -v edn/...
