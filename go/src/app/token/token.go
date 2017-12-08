package token

const (
	// ILLEGAL 何にも識別されなかった文字列
	ILLEGAL = "ILLEGAL"
	// EOF 文字列の終了
	EOF = "EOF"

	// IDENT 変数名
	IDENT = "IDENT"
	// INT integer
	INT = "INT"

	// ASSIGN 代入
	ASSIGN = "="
	// PLUS ＋演算子
	PLUS = "+"
	// MINUS -演算子
	MINUS = "-"
	// BANG !演算子
	BANG = "!"
	// ASTERISK *演算子
	ASTERISK = "*"
	// SLASH /演算子
	SLASH = "/"

	// EQ == 比較演算子
	EQ = "=="
	//NOT_EQ a != 比較演算子
	NOT_EQ = "!=" // nolint

	// LT < 比較演算子
	LT = "<"
	// GT > 比較演算子
	GT = ">"

	// COMMA コンマ
	COMMA = ","
	// SEMICOLON コンマ
	SEMICOLON = ";"

	// LPAREN (
	LPAREN = "("
	// RPAREN )
	RPAREN = ")"
	// LBRACE {
	LBRACE = "{"
	// RBRACE }
	RBRACE = "}"

	// FUNCTION 関数
	FUNCTION = "FUNCTION"
	// LET 変数宣言
	LET = "LET"
	// TRUE 真
	TRUE = "TRUE"
	// FALSE 偽
	FALSE = "FALSE"
	// IF if文
	IF = "IF"
	// ELSE else文
	ELSE = "ELSE"
	// RETURN return
	RETURN = "RETURN"
)

// TokenType 文字列 Asciiのみ対応想定
type TokenType string // nolint

// Token 文字列情報
type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent 特別な意味をもった文字列かチェックする
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

// type LetStatement struct {
// 	Token Token
// 	Name  *Identifier
// 	Value Expression
// }
//
// func (ls *LetStatement) statementNode()       {}
// func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
//
// type Identifier struct {
// 	Token token.Token
// 	Value string
// }
//
// func (i *Identifier) expressionNode()      {}
// func (i *Identifier) TokenLiteral() string { return i.Literal }
