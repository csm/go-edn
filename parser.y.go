//line parser.y:2
package edn

import __yyfmt__ "fmt"

//line parser.y:2
/*
If this file is not parser.y, it was generated from parser.y and
should not be edited directly.
*/

import (
	"github.com/csm/go-edn/types"
)

func init() {
	//yyDebug = 4
}

//line parser.y:19
type yySymType struct {
	yys int
	k   types.Value
	v   types.Value
	raw string
}

const tOpenBracket = 57346
const tCloseBracket = 57347
const tOpenParen = 57348
const tCloseParen = 57349
const tOpenBrace = 57350
const tCloseBrace = 57351
const tOctothorpe = 57352
const tString = 57353
const tKeyword = 57354
const tCharacter = 57355
const tSymbol = 57356
const tWhitespace = 57357
const tNil = 57358
const tTrue = 57359
const tFalse = 57360
const tInteger = 57361
const tBigInteger = 57362
const tFloat = 57363
const tRational = 57364

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"tOpenBracket",
	"tCloseBracket",
	"tOpenParen",
	"tCloseParen",
	"tOpenBrace",
	"tCloseBrace",
	"tOctothorpe",
	"tString",
	"tKeyword",
	"tCharacter",
	"tSymbol",
	"tWhitespace",
	"tNil",
	"tTrue",
	"tFalse",
	"tInteger",
	"tBigInteger",
	"tFloat",
	"tRational",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.y:165

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 45
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 116

var yyAct = [...]int{

	42, 2, 6, 3, 41, 56, 7, 38, 37, 50,
	49, 5, 5, 15, 40, 40, 44, 51, 5, 5,
	28, 4, 39, 22, 21, 20, 19, 45, 18, 43,
	17, 16, 14, 13, 12, 11, 10, 9, 1, 0,
	46, 0, 48, 0, 48, 47, 0, 53, 48, 0,
	0, 57, 58, 54, 0, 0, 0, 59, 60, 24,
	0, 23, 61, 26, 52, 8, 25, 27, 29, 28,
	0, 36, 30, 31, 32, 33, 34, 35, 24, 0,
	23, 0, 26, 0, 8, 25, 27, 29, 28, 0,
	36, 30, 31, 32, 33, 34, 35, 24, 0, 23,
	0, 26, 0, 55, 25, 27, 29, 28, 0, 36,
	30, 31, 32, 33, 34, 35,
}
var yyPact = [...]int{

	-3, -1000, 74, -1000, -3, -1000, -3, -1000, 6, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -3, -3, -1000, -3, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -3,
	-3, 3, -1000, 4, 55, -1000, 93, -4, 74, -1000,
	-1000, -3, -1000, -3, -1000, 7, -1000, -3, -1000, 74,
	-1000, -1000,
}
var yyPgo = [...]int{

	0, 38, 0, 2, 6, 13, 37, 36, 35, 34,
	33, 32, 31, 30, 28, 26, 25, 24, 23, 21,
	3, 4, 17, 16,
}
var yyR1 = [...]int{

	0, 1, 1, 3, 3, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 19,
	20, 20, 2, 2, 21, 21, 13, 13, 18, 12,
	14, 15, 16, 17, 11, 8, 5, 9, 22, 22,
	23, 23, 10, 6, 7,
}
var yyR2 = [...]int{

	0, 3, 1, 1, 4, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 2, 0, 1, 1, 4, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 4, 3, 1,
	1, 3, 3, 3, 3,
}
var yyChk = [...]int{

	-1000, -1, -2, -20, -19, 15, -3, -4, 10, -6,
	-7, -8, -9, -10, -11, -5, -12, -13, -14, -15,
	-16, -17, -18, 6, 4, 11, 8, 12, 14, 13,
	17, 18, 19, 20, 21, 22, 16, -20, -2, -5,
	8, -21, -2, -21, -23, -2, -2, -21, -2, 7,
	5, -22, 9, -3, -4, 10, 9, -3, -2, -20,
	-2, -3,
}
var yyDef = [...]int{

	22, -2, 2, 23, 20, 19, 22, 3, 0, 5,
	6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
	16, 17, 18, 22, 22, 35, 22, 34, 36, 29,
	26, 27, 30, 31, 32, 33, 28, 21, 1, 22,
	22, 22, 24, 22, 0, 40, 0, 22, 0, 43,
	44, 22, 42, 39, 4, 0, 37, 22, 41, 0,
	25, 38,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer, *yySymType) int
	Lookahead() int
}

type yyParserImpl struct {
	lookahead func() int
}

func (p *yyParserImpl) Lookahead() int {
	return p.lookahead()
}

func yyNewParser() yyParser {
	p := &yyParserImpl{
		lookahead: func() int { return -1 },
	}
	return p
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer, result *yySymType) int {
	return yyNewParser().Parse(yylex, result)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer, result *yySymType) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yytoken := -1 // yychar translated into internal numbering
	yyrcvr.lookahead = func() int { return yychar }
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yychar = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar, yytoken = yylex1(yylex, &yylval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yychar = -1
		yytoken = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar, yytoken = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yychar = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:41
		{
			result.v = yyDollar[2].v
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:42
		{
			yylex.Error("empty input")
		}
	case 4:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:47
		{
			yyVAL.v = types.TaggedValue{string(yyDollar[2].v.(types.Symbol)), yyDollar[4].v}
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:80
		{
			yyVAL.v = types.Value(new(types.List))
		}
	case 25:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:81
		{
			yyDollar[1].v.(*types.List).Insert(yyDollar[3].v)
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:87
		{
			yyVAL.v = types.Bool(true)
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:88
		{
			yyVAL.v = types.Bool(false)
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:92
		{
			yyVAL.v = types.Value(nil)
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:129
		{
			set := types.Sequence(types.Set{})
			values := types.Sequence(yyDollar[3].v.(*types.List))
			yyVAL.v = set.Into(values)
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:137
		{
			yyVAL.k = yyDollar[1].v
			yyVAL.v = yyDollar[3].v
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:138
		{
			yylex.Error("Map literal must contain an even number of forms")
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:142
		{
			yyVAL.v = types.Map{}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:144
		{
			yyDollar[1].v.(types.Map)[yyDollar[2].k] = yyDollar[2].v
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:150
		{
			yyVAL = yyDollar[2]
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:154
		{
			yyVAL = yyDollar[2]
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:159
		{
			vec := types.Sequence(types.Vector{})
			values := types.Sequence(yyDollar[2].v.(*types.List))
			yyVAL.v = vec.Into(values)
		}
	}
	goto yystack /* stack new state and value */
}
