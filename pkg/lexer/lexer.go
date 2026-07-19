package lexer

import (
	"strings"

	"github.com/tom96da/sleepingknights/pkg/token"
)

var keywords = map[string]token.TokenType{
	"let":      token.LET,
	"const":    token.CONST,
	"fn":       token.FN,
	"if":       token.IF,
	"else":     token.ELSE,
	"for":      token.FOR,
	"while":    token.WHILE,
	"return":   token.RETURN,
	"break":    token.BREAK,
	"continue": token.CONTINUE,
	"true":     token.TRUE,
	"false":    token.FALSE,
	"int":      token.INT_TYPE,
	"float":    token.FLOAT_TYPE,
	"bool":     token.BOOL_TYPE,
	"string":   token.STRING_TYPE,
	"void":     token.VOID_TYPE,
	"and":      token.AND,
	"or":       token.OR,
	"not":      token.NOT,
}

// Lexer converts source text to lexical tokens while tracking positions.
type Lexer struct {
	input []rune

	position     int
	readPosition int

	ch rune

	line   int
	column int

	nextLine   int
	nextColumn int
}

// New creates a lexer for the provided source input.
func New(input string) *Lexer {
	l := &Lexer{
		input:      []rune(input),
		nextLine:   1,
		nextColumn: 1,
	}
	l.readChar()
	return l
}

// NextToken returns the next token in the input stream.
func (l *Lexer) NextToken() token.Token {
	l.skipInlineWhitespace()

	line := l.line
	column := l.column

	switch l.ch {
	case 0:
		return token.Token{Type: token.EOF, Literal: "", Line: line, Column: column}
	case '\n':
		l.readChar()
		return token.Token{Type: token.NEWLINE, Literal: "\\n", Line: line, Column: column}
	case '+':
		tok := token.Token{Type: token.PLUS, Literal: "+", Line: line, Column: column}
		l.readChar()
		return tok
	case '-':
		tok := token.Token{Type: token.MINUS, Literal: "-", Line: line, Column: column}
		l.readChar()
		return tok
	case '*':
		tok := token.Token{Type: token.STAR, Literal: "*", Line: line, Column: column}
		l.readChar()
		return tok
	case '%':
		tok := token.Token{Type: token.PERCENT, Literal: "%", Line: line, Column: column}
		l.readChar()
		return tok
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			l.readChar()
			return token.Token{Type: token.EQ, Literal: "==", Line: line, Column: column}
		}
		tok := token.Token{Type: token.ASSIGN, Literal: "=", Line: line, Column: column}
		l.readChar()
		return tok
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			l.readChar()
			return token.Token{Type: token.NOT_EQ, Literal: "!=", Line: line, Column: column}
		}
		tok := token.Token{Type: token.ILLEGAL, Literal: "!", Line: line, Column: column}
		l.readChar()
		return tok
	case '<':
		if l.peekChar() == '=' {
			l.readChar()
			l.readChar()
			return token.Token{Type: token.LE, Literal: "<=", Line: line, Column: column}
		}
		tok := token.Token{Type: token.LT, Literal: "<", Line: line, Column: column}
		l.readChar()
		return tok
	case '>':
		if l.peekChar() == '=' {
			l.readChar()
			l.readChar()
			return token.Token{Type: token.GE, Literal: ">=", Line: line, Column: column}
		}
		tok := token.Token{Type: token.GT, Literal: ">", Line: line, Column: column}
		l.readChar()
		return tok
	case '/':
		if l.peekChar() == '/' {
			literal := l.readLineComment()
			return token.Token{Type: token.LINE_COMMENT, Literal: literal, Line: line, Column: column}
		}
		if l.peekChar() == '*' {
			isDoc := l.peekSecondChar() == '*'
			literal, ok := l.readBlockComment()
			if !ok {
				return token.Token{Type: token.ILLEGAL, Literal: literal, Line: line, Column: column}
			}
			if isDoc {
				return token.Token{Type: token.DOC_COMMENT, Literal: literal, Line: line, Column: column}
			}
			return token.Token{Type: token.BLOCK_COMMENT, Literal: literal, Line: line, Column: column}
		}
		tok := token.Token{Type: token.SLASH, Literal: "/", Line: line, Column: column}
		l.readChar()
		return tok
	case '(':
		tok := token.Token{Type: token.LPAREN, Literal: "(", Line: line, Column: column}
		l.readChar()
		return tok
	case ')':
		tok := token.Token{Type: token.RPAREN, Literal: ")", Line: line, Column: column}
		l.readChar()
		return tok
	case '{':
		tok := token.Token{Type: token.LBRACE, Literal: "{", Line: line, Column: column}
		l.readChar()
		return tok
	case '}':
		tok := token.Token{Type: token.RBRACE, Literal: "}", Line: line, Column: column}
		l.readChar()
		return tok
	case ',':
		tok := token.Token{Type: token.COMMA, Literal: ",", Line: line, Column: column}
		l.readChar()
		return tok
	case ':':
		tok := token.Token{Type: token.COLON, Literal: ":", Line: line, Column: column}
		l.readChar()
		return tok
	case ';':
		tok := token.Token{Type: token.SEMICOLON, Literal: ";", Line: line, Column: column}
		l.readChar()
		return tok
	case '"':
		literal, ok := l.readString()
		if !ok {
			return token.Token{Type: token.ILLEGAL, Literal: literal, Line: line, Column: column}
		}
		return token.Token{Type: token.STRING, Literal: literal, Line: line, Column: column}
	}

	if isLetter(l.ch) {
		literal := l.readIdentifier()
		typeValue, ok := keywords[literal]
		if !ok {
			typeValue = token.IDENT
		}
		return token.Token{Type: typeValue, Literal: literal, Line: line, Column: column}
	}

	if isDigit(l.ch) {
		literal, typeValue, ok := l.readNumber()
		if !ok {
			return token.Token{Type: token.ILLEGAL, Literal: literal, Line: line, Column: column}
		}
		return token.Token{Type: typeValue, Literal: literal, Line: line, Column: column}
	}

	illegal := string(l.ch)
	l.readChar()
	return token.Token{Type: token.ILLEGAL, Literal: illegal, Line: line, Column: column}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.position = l.readPosition
		l.ch = 0
		l.line = l.nextLine
		l.column = l.nextColumn
		return
	}

	l.ch = l.input[l.readPosition]
	l.position = l.readPosition
	l.readPosition++

	l.line = l.nextLine
	l.column = l.nextColumn

	if l.ch == '\n' {
		l.nextLine++
		l.nextColumn = 1
	} else {
		l.nextColumn++
	}
}

func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) peekSecondChar() rune {
	if l.readPosition+1 >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition+1]
}

func (l *Lexer) skipInlineWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\f' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	start := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return string(l.input[start:l.position])
}

func (l *Lexer) readNumber() (string, token.TokenType, bool) {
	start := l.position

	if l.ch == '0' && (l.peekChar() == 'x' || l.peekChar() == 'X') {
		l.readChar()
		l.readChar()
		digitStart := l.position
		for isHexDigit(l.ch) {
			l.readChar()
		}
		if l.position == digitStart {
			return string(l.input[start:l.position]), token.ILLEGAL, false
		}
		return string(l.input[start:l.position]), token.INT, true
	}

	if l.ch == '0' && (l.peekChar() == 'b' || l.peekChar() == 'B') {
		l.readChar()
		l.readChar()
		digitStart := l.position
		for l.ch == '0' || l.ch == '1' {
			l.readChar()
		}
		if l.position == digitStart {
			return string(l.input[start:l.position]), token.ILLEGAL, false
		}
		return string(l.input[start:l.position]), token.INT, true
	}

	typeValue := token.INT
	for isDigit(l.ch) {
		l.readChar()
	}

	if l.ch == '.' && isDigit(l.peekChar()) {
		typeValue = token.FLOAT
		l.readChar()
		for isDigit(l.ch) {
			l.readChar()
		}
	}

	if l.ch == 'e' || l.ch == 'E' {
		typeValue = token.FLOAT
		l.readChar()
		if l.ch == '+' || l.ch == '-' {
			l.readChar()
		}
		if !isDigit(l.ch) {
			return string(l.input[start:l.position]), token.ILLEGAL, false
		}
		for isDigit(l.ch) {
			l.readChar()
		}
	}

	return string(l.input[start:l.position]), typeValue, true
}

func (l *Lexer) readString() (string, bool) {
	l.readChar()

	var b strings.Builder
	for {
		switch l.ch {
		case 0, '\n':
			return "unterminated string", false
		case '"':
			l.readChar()
			return b.String(), true
		case '\\':
			l.readChar()
			switch l.ch {
			case 'n':
				b.WriteRune('\n')
			case 't':
				b.WriteRune('\t')
			case 'r':
				b.WriteRune('\r')
			case 'b':
				b.WriteRune('\b')
			case 'f':
				b.WriteRune('\f')
			case '"':
				b.WriteRune('"')
			case '\\':
				b.WriteRune('\\')
			default:
				return "invalid escape sequence", false
			}
			l.readChar()
		default:
			b.WriteRune(l.ch)
			l.readChar()
		}
	}
}

func (l *Lexer) readLineComment() string {
	start := l.position
	for l.ch != 0 && l.ch != '\n' {
		l.readChar()
	}
	return string(l.input[start:l.position])
}

func (l *Lexer) readBlockComment() (string, bool) {
	start := l.position

	// Consume "/*".
	l.readChar()
	l.readChar()

	for {
		if l.ch == 0 {
			return "unterminated block comment", false
		}
		if l.ch == '*' && l.peekChar() == '/' {
			l.readChar()
			l.readChar()
			return string(l.input[start:l.position]), true
		}
		l.readChar()
	}
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isHexDigit(ch rune) bool {
	return isDigit(ch) || (ch >= 'a' && ch <= 'f') || (ch >= 'A' && ch <= 'F')
}
