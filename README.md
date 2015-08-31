# EDN

An [EDN][edn] parser for Go.

## Building

tl;dr:

```
go get github.com/csm/go-edn
go generate github.com/csm/go-edn
```

The first command will fetch the sources, but fail to build; the second command
will generate the lexer (with [nex][nex]) and parser (with [yacc][yacc]).

Tested with go 1.5; 1.4 may work too.

## License

MPL 2.0

[edn]: https://github.com/edn-format/edn
[nex]: http://www-cs-students.stanford.edu/~blynn/nex/
[yacc]: http://golang.org/cmd/yacc/
[gohash]: https://code.google.com/p/gohash/
[gohash interfaces]: https://code.google.com/p/gohash/source/browse/hash/set.go#37
[irc discussion]: https://botbot.me/freenode/go-nuts/msg/3137158/
[go build advice]: http://golang.org/doc/articles/go_command.html#tmp_4
