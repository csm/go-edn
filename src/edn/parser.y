%{
package edn 

/*
If this file is not parser.y, it was generated from parser.y and
should not be edited directly.
*/

import (
	"edn/types"
)

func init() {
	//yyDebug = 4
}

var result = new(yySymType)

%}

%union { 
	k types.Value
	v types.Value
	raw string
} 

%token tOpenBracket tCloseBracket
%token tOpenParen tCloseParen
%token tOpenBrace tCloseBrace
%token tOctothorpe
%token tString tKeyword tCharacter tSymbol
%token tWhitespace
%token tNil
%token tTrue
%token tFalse
%token tInteger
%token tBigInteger
%token tFloat
%token tRational

%% 
 input /* somebody help me not have to do this: */
	: ws✳ value ws✳ { result.v = $2.v }
	| ws✳ { yylex.Error("empty input") }
	;

value
	: untaggedValue
	| tOctothorpe symbol ws✳ untaggedValue { $$.v = types.TaggedValue{string($2.v.(types.Symbol)), $4.v} }
	;

untaggedValue
	: list
	| vector
	| string
	| set
	| map
	| keyword
	| symbol
	| character
	| bool
	| integer
	| bigint
	| float
	| rational
	| nil
	;

ws
	: tWhitespace
	;

ws＋
	: ws
	| ws ws＋

ws✳
	: /* empty */
	| ws＋

values
	: ws✳ { $$.v = types.Value(new(types.List))}
	| values ws✳ value ws✳ {
		$1.v.(*types.List).Insert($3.v)
	  }
	;

bool
	: tTrue { $$.v = types.Bool(true) }
	| tFalse { $$.v = types.Bool(false) }
	;

nil
	: tNil { $$.v = types.Value(nil) }
	;

character
	: tCharacter
	;

integer
	: tInteger
	;

bigint
	: tBigInteger
	;

float
	: tFloat
	;

rational
	: tRational
	;

keyword
	: tKeyword
	;

string
	: tString
	;

symbol
	: tSymbol
	;

set
	: tOctothorpe tOpenBrace values tCloseBrace
	  { 
	  	set := types.Sequence(types.Set{})
	  	values := types.Sequence($3.v.(*types.List))
	  	$$.v = set.Into(values)
	  }
	;

key_value
	: value ws＋ value { $$.k = $1.v; $$.v = $3.v }
	| value { yylex.Error("Map literal must contain an even number of forms") }
	;

key_values
	: ws✳ { $$.v = types.Map{} }
	| key_values key_value ws✳
	  {
	  	$1.v.(types.Map)[$2.k] = $2.v
	  }
	;

map
	: tOpenBrace key_values tCloseBrace { $$ = $2 }
	;

list
	: tOpenParen values tCloseParen { $$ = $2 }
	;

vector
	: tOpenBracket values tCloseBracket 
	  {
	  	vec := types.Sequence(types.Vector{})
	  	values := types.Sequence($2.v.(*types.List))
		$$.v = vec.Into(values)
	  }
	;
%%
