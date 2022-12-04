package lexer

type Lexer struct {
	input        string
	charStorage  []byte
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
	tempCh       byte // char to peek at
}
