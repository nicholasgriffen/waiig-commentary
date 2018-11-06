package lexer
import "interpreter/token"
//Lexer takes source code and returns tokens 
type Lexer struct {
	input string 
	position, nextPosition int
	character byte
}
//New takes string and returns initialized lexer
func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.nextCharacter()
	return lexer
}

func (lexer *Lexer) nextCharacter() {
	if lexer.nextPosition >= len(lexer.input) {
		//ASCII for NUL
		lexer.character = 0
	} else {
		lexer.character = lexer.input[lexer.nextPosition]
	}

	lexer.position = lexer.nextPosition
	lexer.nextPosition ++
}
//NextToken creates token and steps position
func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token 
	lexer.skipWhitespace()

	switch lexer.character {
	case '=':
		tok = newToken(token.ASSIGN, lexer.character)
	case '+':
		tok = newToken(token.PLUS, lexer.character)
	case ',':
		tok = newToken(token.COMMA, lexer.character)
	case ';':
		tok = newToken(token.SEMICOLON, lexer.character)
	case '(':
		tok = newToken(token.LROUND, lexer.character)
	case ')':
		tok = newToken(token.RROUND, lexer.character)
	case '{':
		tok = newToken(token.LCURLY, lexer.character)
	case '}':
		tok = newToken(token.RCURLY, lexer.character)
	case '1':
		tok = newToken(token.INT, lexer.character)
	case '2':
		tok = newToken(token.INT, lexer.character)
	case '3':
		tok = newToken(token.INT, lexer.character)
	case '4':
		tok = newToken(token.INT, lexer.character)
	case '5':
		tok = newToken(token.INT, lexer.character)
	case '6':
		tok = newToken(token.INT, lexer.character)
	case '7':
		tok = newToken(token.INT, lexer.character)
	case '8':
		tok = newToken(token.INT, lexer.character)
	case '9':
		tok = newToken(token.INT, lexer.character)
	case '0':
		tok.Literal = ""
		tok.Type = token.EOF
	default: 
		if isLetter(lexer.character) {
			tok.Literal = lexer.checkIdentifier()
			tok.Type = token.FindIdent(tok.Literal)
			return tok
		} 
		tok = newToken(token.ILLEGAL, lexer.character)
	}

	lexer.nextCharacter()
	return tok
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.character == ' ' || lexer.character == '\n' || lexer.character == '\t' || lexer.character == '\r' {
		lexer.nextCharacter()
	}
}

func newToken(tokenType token.TokenType, character byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(character)}
}

func (lexer *Lexer) checkIdentifier() string {
	position := lexer.position 
	for isLetter(lexer.character) {
		lexer.nextCharacter()
	}
	return lexer.input[position:lexer.position]
}

func isLetter(character byte) bool {
	return 'a' <= character && character <= 'z' || 'A' <= character && character <= 'Z' || character == '_' 
}
//PRODUCTION wants Unicode UTF-8 support