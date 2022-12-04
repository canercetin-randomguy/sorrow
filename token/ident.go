package token

func LookupIdent(ident string) TokenType {
	// is this a keyword?
	if tok, ok := Keywords[ident]; ok {
		return tok
	}
	// if not, it is a variable.
	return VARIABLE
}
