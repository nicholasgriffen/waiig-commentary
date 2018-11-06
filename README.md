# Quick Start
## Install
```bash
git clone git@github.com:nicholasgriffen/waiig-commentary.git
```
## Start repl 
```bash
cd waiig-commentary
go run .
```
# repl
--
    import "interpreter/repl"


## Usage

#### func  Start

```go
func Start(in io.Reader, out io.Writer)
```
Start reads lines in and writes tokens out
# lexer
--
    import "interpreter/lexer"


## Usage

#### type Lexer

```go
type Lexer struct {
}
```

Lexer takes source code and returns tokens

#### func  New

```go
func New(input string) Lexer
```
New takes string and returns initialized lexer

#### func (*Lexer) NextToken

```go
func (lexer *Lexer) NextToken() token.Token
```
NextToken creates token and steps position
# token
--
    import "interpreter/token"

BUG(nicholasgriffen): Tokens lack filenames, linenumbers

## Usage

```go
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
	EOF = ""
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
```

#### type Token

```go
type Token struct {
	Type    TokenType
	Literal string
}
```

Token keywords, symbols

#### type TokenType

```go
type TokenType string
```

TokenType matters downstream

#### func  FindIdent

```go
func FindIdent(ident string) TokenType
```
FindIdent takes string and looks up TokenType
