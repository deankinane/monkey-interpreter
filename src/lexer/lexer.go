package lexer

import "github.com/deankinane/monkey-interpreter/src/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(source string) *Lexer {
	l := &Lexer{input: source}
	l.readChar()
	return l
}

func NewToken(tokenType token.TokenType, literal byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(literal),
	}
}

func (l *Lexer) NextToken() token.Token {
	tok := token.Token{}

	tok.Literal = string(l.ch)

	switch l.ch {
	case '=':
		tok.Type = token.ASSIGN
	case '+':
		tok.Type = token.PLUS
	case ',':
		tok.Type = token.COMMA
	case '(':
		tok.Type = token.LPAREN
	case ')':
		tok.Type = token.RPAREN
	case '{':
		tok.Type = token.LBRACE
	case '}':
		tok.Type = token.RBRACE
	case ';':
		tok.Type = token.SEMICOLON
	case 0:
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.GetIdent(tok.Literal)
			return tok
		} else {
			tok.Type = token.ILLEGAL
		}
	}

	l.readChar()

	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readIdentifier() string {
	pos := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.position]
}
