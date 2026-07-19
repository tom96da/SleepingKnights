package token

// TokenType identifies the kind of lexical token.
type TokenType string

const (
	// Special tokens

	ILLEGAL TokenType = "ILLEGAL" // Token for invalid or unrecognized input.
	EOF     TokenType = "EOF"     // Token marking the end of the input stream.

	// Literals and Identifiers

	IDENT  TokenType = "IDENT"  // Token for an identifier name.
	INT    TokenType = "INT"    // Token for an integer literal.
	FLOAT  TokenType = "FLOAT"  // Token for a floating-point literal.
	STRING TokenType = "STRING" // Token for a string literal.

	// Keywords

	LET      TokenType = "LET"      // Token for the `let` variable declaration keyword.
	CONST    TokenType = "CONST"    // Token for the `const` immutable binding declaration keyword.
	FN       TokenType = "FN"       // Token for the `fn` function declaration keyword.
	IF       TokenType = "IF"       // Token for the `if` conditional branch keyword.
	ELSE     TokenType = "ELSE"     // Token for the `else` alternate conditional branch keyword.
	FOR      TokenType = "FOR"      // Token for the `for` counted or general loop keyword.
	WHILE    TokenType = "WHILE"    // Token for the `while` condition-controlled loop keyword.
	RETURN   TokenType = "RETURN"   // Token for the `return` function return keyword.
	BREAK    TokenType = "BREAK"    // Token for the `break` loop termination keyword.
	CONTINUE TokenType = "CONTINUE" // Token for the `continue` loop continuation keyword.

	// Boolean keywords

	TRUE  TokenType = "TRUE"  // Token for the `true` boolean literal.
	FALSE TokenType = "FALSE" // Token for the `false` boolean literal.

	// Type keywords

	INT_TYPE    TokenType = "INT_TYPE"    // Token for the `int` integer type keyword.
	FLOAT_TYPE  TokenType = "FLOAT_TYPE"  // Token for the `float` floating-point type keyword.
	BOOL_TYPE   TokenType = "BOOL_TYPE"   // Token for the `bool` boolean type keyword.
	STRING_TYPE TokenType = "STRING_TYPE" // Token for the `string` string type keyword.
	VOID_TYPE   TokenType = "VOID_TYPE"   // Token for the `void` type keyword.

	// Logical operators

	AND TokenType = "AND" // Token for the `and` logical conjunction operator.
	OR  TokenType = "OR"  // Token for the `or` logical disjunction operator.
	NOT TokenType = "NOT" // Token for the `not` logical negation operator.

	// Arithmetic operators

	PLUS    TokenType = "PLUS"    // Token for the `+` addition operator.
	MINUS   TokenType = "MINUS"   // Token for the `-` subtraction or unary negation operator.
	STAR    TokenType = "STAR"    // Token for the `*` multiplication operator.
	SLASH   TokenType = "SLASH"   // Token for the `/` division operator.
	PERCENT TokenType = "PERCENT" // Token for the `%` remainder operator.

	// Comparison operators

	EQ     TokenType = "EQ"     // Token for the `==` equality comparison operator.
	NOT_EQ TokenType = "NOT_EQ" // Token for the `!=` inequality comparison operator.
	LT     TokenType = "LT"     // Token for the `<` less-than comparison operator.
	LE     TokenType = "LE"     // Token for the `<=` less-than-or-equal comparison operator.
	GT     TokenType = "GT"     // Token for the `>` greater-than comparison operator.
	GE     TokenType = "GE"     // Token for the `>=` greater-than-or-equal comparison operator.

	// Assignment and delimiters

	ASSIGN    TokenType = "ASSIGN"    // Token for the `=` assignment operator.
	LPAREN    TokenType = "LPAREN"    // Token for the `(` left parenthesis delimiter.
	RPAREN    TokenType = "RPAREN"    // Token for the `)` right parenthesis delimiter.
	LBRACE    TokenType = "LBRACE"    // Token for the `{` left brace delimiter.
	RBRACE    TokenType = "RBRACE"    // Token for the `}` right brace delimiter.
	COMMA     TokenType = "COMMA"     // Token for the `,` comma separator.
	COLON     TokenType = "COLON"     // Token for the `:` colon separator.
	SEMICOLON TokenType = "SEMICOLON" // Token for the `;` semicolon separator.
	NEWLINE   TokenType = "NEWLINE"   // Token for the `\n` newline statement terminator.

	// Comments

	LINE_COMMENT  TokenType = "LINE_COMMENT"  // Token for a `//` line comment.
	BLOCK_COMMENT TokenType = "BLOCK_COMMENT" // Token for a `/* */` block comment.
	DOC_COMMENT   TokenType = "DOC_COMMENT"   // Token for a `/** */` documentation comment.
)

// Token stores a lexical token and its source position.
type Token struct {
	Type    TokenType
	Literal string
	Line    int // Line number (1-indexed)
	Column  int // Column number (1-indexed)
}
