package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
	"if": IF,
	"else": ELSE,
	"true": TRUE,
	"false": FALSE,
	"return": RETURN,
}

func LookupIdentifier(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENTIFIER
}
