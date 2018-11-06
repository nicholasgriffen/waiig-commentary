package token
//TokenType matters downstream
type TokenType string

const ( 
	//ILLEGAL unknown type
	ILLEGAL = "ILLEGAL"
	//IDENT identifier
	IDENT = "IDENT"
	//FUNCTION function declaration
	FUNCTION = "FUNCTION"
	//RETURN end function execution with return value
	RETURN = "RETURN"
	//LET identifier naming
	LET = "LET"
	//TRUE boolean false
	TRUE = "TRUE"
	//FALSE boolean false
	FALSE = "FALSE"
	//IF if control flow
	IF = "IF"
	//ELSE else control flow
	ELSE = "ELSE"
	//IS comparison
	IS = "=="
	//ISNOT comparison
	ISNOT = "!="
	//EOF end of file
	EOF	= ""
	//INT integer
	INT = "INT"
	//ASSIGN identifier assignment
	ASSIGN = "="
	//PLUS sum
	PLUS = "+"
	//MINUS subtract
	MINUS = "-"
	//STAR asterisk
	STAR = "*"
	//FSLASH forward slash
	FSLASH = "/"
	//COMMA expression separator{}
	COMMA = ","
	//SEMICOLON statement terminator
	SEMICOLON = ";"
	//BANG exclamation point
	BANG = "!"
	//LTHAN open angle
	LTHAN = "<"
	//RTHAN close angle
	RTHAN = ">"
	//LROUND open grouping
	LROUND = "("
	//RROUND close grouping
	RROUND = ")"
	//LCURLY open block
	LCURLY = "{"
	//RCURLY close block
	RCURLY = "}"	
)

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}
//FindIdent takes string and looks up TokenType
func FindIdent(ident string) TokenType {
	if token, ok := keywords[ident]; ok {
		return token
	}
	return IDENT
}

//Token keywords, symbols
type Token struct {
	Type	TokenType
	Literal string
}


//PRODUCTION wants filenames linenumbers