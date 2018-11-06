//BUG(nicholasgriffen): Lexers lack Unicode, UTF-8 support
package lexer
import "interpreter/token"
//Lexer takes source code and returns tokens 
type Lexer struct {
	input string 
	position, nextPosition int
	character byte
}
//New takes string and returns initialized lexer
func New(input string) Lexer {
	lexer := Lexer{input: input}
	lexer.setNextCharacter()
	return lexer
}

func (lexer *Lexer) setNextCharacter() {
	if lexer.nextPosition >= len(lexer.input) {
		//ASCII for NUL
		lexer.character = 0
	} else {
		lexer.character = lexer.input[lexer.nextPosition]
	}

	lexer.position = lexer.nextPosition
	lexer.nextPosition ++
}

func (lexer *Lexer) getNextCharacter() byte {
	if lexer.nextPosition >= len(lexer.input) {
		return 0
	}	
	return lexer.input[lexer.nextPosition]
}
//NextToken creates token and steps position
func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token 
	lexer.skipWhitespace()

	switch lexer.character {
	case '=':
		if lexer.getNextCharacter() == '=' {
			previousCharacter := lexer.character
			lexer.setNextCharacter()
			tok = token.Token{Type: token.IS, Literal: string(previousCharacter) + string(lexer.character)}
		} else {
			tok = newToken(token.ASSIGN, lexer.character)
		}
	case '+':
		tok = newToken(token.PLUS, lexer.character)
	case '-':
		tok = newToken(token.MINUS, lexer.character)
	case '*':
		tok = newToken(token.STAR, lexer.character)
	case '/':
		tok = newToken(token.FSLASH, lexer.character)
	case ',':
		tok = newToken(token.COMMA, lexer.character)
	case ';':
		tok = newToken(token.SEMICOLON, lexer.character)
	case '!':
		if lexer.getNextCharacter() == '=' {
			previousCharacter := lexer.character
			lexer.setNextCharacter()
			tok = token.Token{Type: token.ISNOT, Literal: string(previousCharacter) + string(lexer.character)}
		} else {
			tok = newToken(token.BANG, lexer.character)
		}
	case '<':
		tok = newToken(token.LTHAN, lexer.character)
	case '>':
		tok = newToken(token.RTHAN, lexer.character)
	case '(':
		tok = newToken(token.LROUND, lexer.character)
	case ')':
		tok = newToken(token.RROUND, lexer.character)
	case '{':
		tok = newToken(token.LCURLY, lexer.character)
	case '}':
		tok = newToken(token.RCURLY, lexer.character)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default: 
		if isLetter(lexer.character) {
			tok.Literal = lexer.parseIdentifier() // 1
			tok.Type = token.FindIdent(tok.Literal) // 2
			return tok
		} else if isDigit(lexer.character) {
			tok.Literal = lexer.parseNumber()
			tok.Type = token.INT 
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lexer.character)
		}
	}

	lexer.setNextCharacter()
	return tok
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.character == ' ' || lexer.character == '\n' || lexer.character == '\t' || lexer.character == '\r' {
		lexer.setNextCharacter()
	}
}

func newToken(tokenType token.TokenType, character byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(character)}
}

func isLetter(character byte) bool {
	return 'a' <= character && character <= 'z' || 'A' <= character && character <= 'Z' || character == '_' 
}

func isDigit(character byte) bool {
	return '0' <= character && character <= '9'
}

func (lexer *Lexer) parseIdentifier() string {
	firstCharacter := lexer.position 
	for isLetter(lexer.character) {
		lexer.setNextCharacter()
	}
	return lexer.input[firstCharacter:lexer.position]
}

func (lexer *Lexer) parseNumber() string {
	position := lexer.position 
	for isDigit(lexer.character) {
		lexer.setNextCharacter()
	}
	return lexer.input[position:lexer.position]
}