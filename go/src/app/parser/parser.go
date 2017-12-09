package parser

import (
	"app/ast"
	"app/lexer"
	"app/token"
	"fmt"
)

// Parser 解析
type Parser struct {
	// lexerへのポインタ
	l *lexer.Lexer

	// 現在のトークン情報
	curToken token.Token
	// 現在のトークンだけでは不十分の場合に次のトークン見る時に使う
	// 例: 5ときたら、ここで終了なのか算術式が来るのか？など
	peekToken token.Token
	// 発生したエラー
	errors []string
}

// New コンストラクタ
// 最初の段階で最初のtokenと次のtokenが必要なので、内部で2回 nextTokenが実行される
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// nextToken 現在の位置の次のtokenをcurTokenに入れ、次のpeekTokenにさらに次のtokenを入れる
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram EOFになるまで解析する
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

// parseStatement それぞれの構文チェックをする
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

// parseLetStatement 変数宣言の構文チェックをする
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	// let の 次に変数名に当たるものが来ているか
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// let hoge の次に = が来ているか
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: 現状は値そのものはチェックしない

	// let hoge = 1 などの次に ; 来ているか調べる
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()

	// TODO 型チェックは行わない
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// curTokenIs 引数のtokenTypeと現在のcurTokenのtypeが同じか調べる
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// peekTokenIs 引数のtokenTypeと現在のpeekTokenのtypeが同じか調べる
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// expectPeek 引数のtokenTypeと現在のpeekTokenのtypeが同じか調べる
// もし同じだった場合、次のtokenを取得する
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}
