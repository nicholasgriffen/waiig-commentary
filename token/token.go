package token
//Type matters downstream
type Type string
//Token keywords, symbols
type Token struct {
	Type	Type
	Literal string
}

const ( 
	//ILLEGAL unknown type
	ILLEGAL = "ILLEGAL"
	//EOF end of file
	EOF	= "EOF"
	//IDENT identifier
	IDENT = "IDENT"
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
	//FUNCTION function declaration
	FUNCTION = "FUNCTION"
	//LET identifier naming
	LET = "LET"	
)

//PRODUCTION wants filenames linenumbers