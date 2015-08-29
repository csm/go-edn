s:Parse(yylex yyLexer) int:Parse(yylex yyLexer, result *yySymType) int:
s:yyNewParser().Parse(yylex):yyNewParser().Parse(yylex, result):
s:Parse(yyLexer) int:Parse(yyLexer, *yySymType) int: