package ast

import (
	"app/token"
	"bytes"
)

// Node AST Node interface
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement
type Statement interface {
	Node
	statementNode()
}

// Expression
type Expression interface {
	Node
	expressionNode()
}

// Program
type Program struct {
	Statements []Statement
}

// String
func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// TokenLiteral 関連付けられている文字列を取得する
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement バインディングの識別子を保持する名前と、値を生成する式の値
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

// statementNode
func (ls *LetStatement) statementNode() {}

// TokenLiteral 関連付けられている文字列を取得する
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// String
func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

// Identifier
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

// expressionNode
func (i *Identifier) expressionNode() {}

// TokenLiteral
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// String
func (i *Identifier) String() string { return i.Value }

// ReturnStatement
type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

// statementNode
func (rs *ReturnStatement) statementNode() {}

// TokenLiteral
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// String
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// ExpressionStatement
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

// statementNode
func (es *ExpressionStatement) statementNode() {}

// TokenLiteral
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

// String
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

type PrefixExpression struct {
	Token    token.Token
	Operator string
	// 演算子の右側の式
	Right Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (oe *InfixExpression) expressionNode()      {}
func (oe *InfixExpression) TokenLiteral() string { return oe.Token.Literal }
func (oe *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")
	return out.String()
}
