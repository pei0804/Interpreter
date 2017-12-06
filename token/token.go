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
)

// TokenType 文字列 Asciiのみ対応想定
type TokenType string

// Token 文字列情報
type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdent 特別な意味をもった文字列かチェックする
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
