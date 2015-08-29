//line parser.y:2
package edn

import __yyfmt__ "fmt"

//line parser.y:2
/*
If this file is not parser.y, it was generated from parser.y and
should not be edited directly.
*/

import "types"

func init() {
	//yyDebug = 4
}

//line parser.y:16
type yySymType struct {
	yys int
	k   types.Value
	v   types.Value
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
const tWhitespace = 57356
const tNil = 57357
const tTrue = 57358
const tFalse = 57359

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
	"tWhitespace",
	"tNil",
	"tTrue",
	"tFalse",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.y:126

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 40,
	9, 26,
	-2, 14,
}

const yyNprod = 32
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 61

var yyAct = [...]int{

	29, 2, 6, 5, 28, 31, 17, 27, 16, 32,
	20, 39, 19, 18, 21, 22, 38, 25, 23, 24,
	42, 33, 30, 36, 4, 5, 35, 15, 14, 34,
	13, 34, 5, 5, 12, 40, 37, 41, 34, 43,
	3, 44, 45, 11, 17, 26, 16, 46, 20, 10,
	19, 18, 21, 22, 9, 25, 23, 24, 8, 7,
	1,
}
var yyPact = [...]int{

	-11, -1000, 40, -1000, -11, -1000, -11, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -11, -11, -1000, -3,
	-11, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 19, -1000,
	18, -11, 2, -1000, 40, -1000, -1000, 11, -11, -1000,
	-11, -11, -1000, -1000, 40, -1000, -1000,
}
var yyPgo = [...]int{

	0, 60, 0, 2, 59, 58, 54, 49, 43, 34,
	30, 28, 27, 24, 40, 4, 16, 9,
}
var yyR1 = [...]int{

	0, 1, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 13, 14, 14, 2, 2, 15, 15, 11, 11,
	12, 10, 9, 6, 7, 16, 16, 17, 17, 8,
	4, 5,
}
var yyR2 = [...]int{

	0, 3, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 2, 0, 1, 1, 4, 1, 1,
	1, 1, 1, 1, 4, 3, 1, 1, 3, 3,
	3, 3,
}
var yyChk = [...]int{

	-1000, -1, -2, -14, -13, 14, -3, -4, -5, -6,
	-7, -8, -9, -10, -11, -12, 6, 4, 11, 10,
	8, 12, 13, 16, 17, 15, -14, -2, -15, -2,
	-15, 8, -17, -2, -2, 7, 5, -15, -16, 9,
	-3, -3, 9, -2, -2, -2, -3,
}
var yyDef = [...]int{

	14, -2, 0, 15, 12, 11, 14, 2, 3, 4,
	5, 6, 7, 8, 9, 10, 14, 14, 23, 0,
	14, 22, 21, 18, 19, 20, 13, 1, 14, 16,
	14, 14, 0, 27, 0, 30, 31, 14, 14, 29,
	-2, 14, 24, 28, 0, 17, 25,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17,
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
	Parse(yyLexer) int
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
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
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
		//line parser.y:33
		{
			result.v = yyDollar[2].v
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:61
		{
			yyVAL.v = types.Value(new(types.List))
		}
	case 17:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:62
		{
			yyDollar[1].v.(*types.List).Insert(yyDollar[3].v)
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:73
		{
			yyVAL.v = types.Value(nil)
		}
	case 24:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:90
		{
			set := types.Sequence(types.Set{})
			values := types.Sequence(yyDollar[3].v.(*types.List))
			yyVAL.v = set.Into(values)
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:98
		{
			yyVAL.k = yyDollar[1].v
			yyVAL.v = yyDollar[3].v
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:99
		{
			yylex.Error("Map literal must contain an even number of forms")
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:103
		{
			yyVAL.v = types.Map{}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:105
		{
			yyDollar[1].v.(types.Map)[yyDollar[2].k] = yyDollar[2].v
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:111
		{
			yyVAL = yyDollar[2]
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:115
		{
			yyVAL = yyDollar[2]
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:120
		{
			vec := types.Sequence(types.Vector{})
			values := types.Sequence(yyDollar[2].v.(*types.List))
			yyVAL.v = vec.Into(values)
		}
	}
	goto yystack /* stack new state and value */
}
