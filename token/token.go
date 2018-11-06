package token
//Type matters downstream
type TokenType string

const ( 
	//ILLEGAL unknown type
	ILLEGAL = "ILLEGAL"
	//IDENT identifier
	IDENT = "IDENT"
	//FUNCTION function declaration
	FUNCTION = "FUNCTION"
	//LET identifier naming
	LET = "LET"
	//EOF end of file
	EOF	= ""
	//INT integer
	INT = "INT"
	//ASSIGN identifier assignment
	ASSIGN = "="
	//PLUS sum
	PLUS = "+"
	//COMMA expression separator
	COMMA = ","
	//SEMICOLON statement terminator
	SEMICOLON = ";"
	//LROUND open grouping
	LROUND = "("
	//RROUND close grouping
	RROUND = ")"
	//LCURLY open block
	LCURLY = "{"
	//RCURLY close block
	RCURLY = "}"	
)

//Token keywords, symbols
type Token struct {
	Type	TokenType
	Literal string
}


//PRODUCTION wants filenames linenumbers