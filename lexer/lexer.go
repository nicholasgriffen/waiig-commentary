package lexer
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
//
func (lexer *Lexer) NextToken() token.Token {
	var token token.Token 

	switch lexer.character {
	case '=':
		token = newToken(token.ASSIGN, lexer.character)
	case '+':
		token = newToken(token.PLUS, lexer.character)
	}
}

//PRODUCTION wants Unicode UTF-8 support