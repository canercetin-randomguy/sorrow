package lexer

import (
	"sorrow/token"
)

func New(input string) *Lexer {
	l := &Lexer{input: input}
	// Initialize the lexer by reading the first character
	l.readChar()
	return l
}
func (l *Lexer) CheckSpecialCharacter(character byte) (token.Token, bool) {
	switch character {
	case '=':
		l.PeekChar()
		if l.tempCh == '=' {
			return NewTokenMultipleLetter(token.EQ, []byte("==")), true
		}
		return NewToken(token.ASSIGN, '='), true
	case ';':
		return NewToken(token.SEMICOLON, ';'), true
	case '(':
		return NewToken(token.LPAREN, '('), true
	case ')':
		return NewToken(token.RPAREN, ')'), true
	case ',':
		return NewToken(token.COMMA, ','), true
	case '+':
		return NewToken(token.PLUS, '+'), true
	case '{':
		return NewToken(token.LBRACE, '{'), true
	case '}':
		return NewToken(token.RBRACE, '}'), true
	default:
		return token.Token{}, false
	}
}
func (l *Lexer) NextToken() token.Token {
	// Declare a new dummy token variable
	var tok token.Token
	if _, ok := l.CheckSpecialCharacter(l.ch); ok {
		tok = NewToken(token.TokenType(l.ch), l.ch)
	} else {
		for {
			if IsLetter(l.ch) { // is it letter?
				l.charStorage = append(l.charStorage, l.ch)
				l.PeekChar()
				tok, ok = l.CheckSpecialCharacter(l.tempCh)
				if ok { // is the next character a special character?
					tok = NewTokenMultipleLetter(token.LookupIdent(string(l.charStorage)), l.charStorage)
					break
				}
				if len(tok.Type) > 0 || l.tempCh == ' ' { // if it is special character or a space, return the token and break.
					tok = NewTokenMultipleLetter(token.LookupIdent(string(l.charStorage)), l.charStorage)
					break
				}
				l.readChar() // if none of the above, read next letter.
			} else if l.ch == ' ' { // if not, is it space?
				tok = NewToken(token.SPACE, l.ch)
				break
			} else if IsNumber(l.ch) { // if not, is it number?
				tok = NewToken(token.INT, l.ch)
				break
			} else if tempTok, ok := l.CheckSpecialCharacter(l.ch); ok { // if not, then it may be a special character.
				if tempTok.Type == token.SEMICOLON {
					tok = NewToken(token.EOF, l.ch)
				}
				break
			} else {
				return token.Token{
					Type:    token.EOF,
					Literal: ",",
				}
			}
		}
	}
	l.readChar()
	l.charStorage = nil
	return tok
}
func IsLetter(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char == '_')
}
func IsNumber(char byte) bool {
	return char >= '0' && char <= '9'
}
func NewToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}
func NewTokenMultipleLetter(tokenType token.TokenType, ch []byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}
func (l *Lexer) readChar() {
	// if reading position is equal or higher than the length of the input string
	if l.readPosition >= len(l.input) {
		// set the character to 0
		l.ch = 0
	} else {
		// else, set the character to the byte at the current reading position
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
func (l *Lexer) PeekChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.tempCh = l.input[l.readPosition]
	}
}
