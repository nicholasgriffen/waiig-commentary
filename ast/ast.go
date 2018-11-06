package ast
//Node provides TokenLiteral() Token{}[Literal]
type Node interface {
	TokenLiteral() string
}
//Statement is a wrapper
type Statement interface {
	Node
	statementNode()
}
//Expression is a wrapper
type Expression interface {
	Node
	expressionNode()
}