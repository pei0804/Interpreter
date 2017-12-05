# Writing an interpreter in Go

## 1-3 test

```go
> github.com/pei0804/interpreter/lexer.(*Lexer).readChar() ./lexer.go:27 (PC: 0x10f194e)
    22:         return l
    23: }
    24:
    25: func (l *Lexer) readChar() {
    26:         // inputされた文字が最後まで参照したかをチェックする
=>  27:         if l.readPosition >= len(l.input) {
    28:                 l.ch = 0
    29:         } else {
    30:                 l.ch = l.input[l.readPosition]
    31:         }
    32:         l.position = l.readPosition
(dlv) p l
*github.com/pei0804/interpreter/lexer.Lexer {
        input: "=+(){},;",
        position: 0,
        readPosition: 0,
        ch: 0,}
(dlv) s
> github.com/pei0804/interpreter/lexer.(*Lexer).readChar() ./lexer.go:30 (PC: 0x10f19ba)
    25: func (l *Lexer) readChar() {
    26:         // inputされた文字が最後まで参照したかをチェックする
    27:         if l.readPosition >= len(l.input) {
    28:                 l.ch = 0
    29:         } else {
                        // input[0]には、=が入ってる
                        // l.chにbyteデータとして格納される
=>  30:                 l.ch = l.input[l.readPosition]
    31:         }
    32:         l.position = l.readPosition
    33:         l.readPosition += 1
    34: }
    35:
(dlv) p l
*github.com/pei0804/interpreter/lexer.Lexer {
        input: "=+(){},;",
        position: 0,
        readPosition: 0,
        ch: 0,}
(dlv) s
> github.com/pei0804/interpreter/lexer.(*Lexer).readChar() ./lexer.go:27 (PC: 0x10f19e8)
    22:         return l
    23: }
    24:
    25: func (l *Lexer) readChar() {
    26:         // inputされた文字が最後まで参照したかをチェックする
=>  27:         if l.readPosition >= len(l.input) {
    28:                 l.ch = 0
    29:         } else {
    30:                 l.ch = l.input[l.readPosition]
    31:         }
    32:         l.position = l.readPosition
(dlv) p l
*github.com/pei0804/interpreter/lexer.Lexer {
        input: "=+(){},;",
        position: 0,
        readPosition: 0,
        ch: 61,}
(dlv)
Warning: listing may not match stale executable
    27:         if l.readPosition >= len(l.input) {
    28:                 l.ch = 0
    29:         } else {
    30:                 l.ch = l.input[l.readPosition]
    31:         }
=>  32:         l.position = l.readPosition
    33:         l.readPosition += 1
    34: }
    35:
    36: func (l *Lexer) NextToken() token.Token {
    37:         var tok token.Token
(dlv) p l
*github.com/pei0804/interpreter/lexer.Lexer {
        input: "=+(){},;",
        position: 0,
        readPosition: 0,
        ch: 61,}
(dlv) s
> github.com/pei0804/interpreter/lexer.(*Lexer).readChar() ./lexer.go:33 (PC: 0x10f1993)
Warning: listing may not match stale executable
    28:                 l.ch = 0
    29:         } else {
    30:                 l.ch = l.input[l.readPosition]
    31:         }
    32:         l.position = l.readPosition
=>  33:         l.readPosition += 1
    34: }
    35:
    36: func (l *Lexer) NextToken() token.Token {
    37:         var tok token.Token
    38:
(dlv) p l
> github.com/pei0804/interpreter/lexer.New() ./lexer.go:22 (PC: 0x10f1908)
Warning: listing may not match stale executable
    17: }
    18:
    19: func New(input string) *Lexer {
    20:         l := &Lexer{input: input}
    21:         l.readChar()
=>  22:         return l
    23: }
    24:
    25: func (l *Lexer) readChar() {
    26:         // inputされた文字が最後まで参照したかをチェックする
    27:         if l.readPosition >= len(l.input) {
(dlv) p l
*github.com/pei0804/interpreter/lexer.Lexer {
        input: "=+(){},;",
        position: 0,
        // readPositionがインクリメントされている
        readPosition: 1,
        ch: 61,}
> github.com/pei0804/interpreter/lexer.TestNextToken() ./lexer_test.go:28 (PC: 0x10f1fe6)
    23:                 {token.COMMA, ","},
    24:                 {token.SEMICOLON, ";"},
    25:                 {token.EOF, ""},
    26:         }
    27:
=>  28:         l := New(input)
    29:
    30:         for i, tt := range tests {
    31:                 tok := l.NextToken()
    32:                 log.Printf("Literal: %s Type: %s", tok.Literal, tok.Type)
    33:
> github.com/pei0804/interpreter/lexer.TestNextToken() ./lexer_test.go:31 (PC: 0x10f20d9)
    26:         }
    27:
    28:         l := New(input)
    29:
    30:         for i, tt := range tests {
                        // 次のtokenを取得する
=>  31:                 tok := l.NextToken()
    32:                 log.Printf("Literal: %s Type: %s", tok.Literal, tok.Type)
    33:
    34:                 if tok.Type != tt.expectedType {
    35:                         t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
    36:                                 i, tt.expectedType, tok.Type)
> github.com/pei0804/interpreter/lexer.(*Lexer).NextToken() ./lexer.go:41 (PC: 0x10f1d04)
Warning: listing may not match stale executable
    36: func (l *Lexer) NextToken() token.Token {
    37:         var tok token.Token
    38:
    39:         switch l.ch {
                // case = とマッチする
    40:         case '=':
=>  41:                 tok = newToken(token.ASSIGN, l.ch)
    42:         case ';':
    43:                 tok = newToken(token.SEMICOLON, l.ch)
    44:         case '(':
    45:                 tok = newToken(token.LPAREN, l.ch)
    46:         case ')':
// トークンに以下の内容が入る
github.com/pei0804/interpreter/token.Token {Type: "=", Literal: "="}

> github.com/pei0804/interpreter/lexer.(*Lexer).NextToken() ./lexer.go:60 (PC: 0x10f1ac4)
Warning: listing may not match stale executable
    55:                 tok = newToken(token.RBRACE, l.ch)
    56:         case 0:
    57:                 tok.Literal = ""
    58:                 tok.Type = token.EOF
    59:         }
                // 取得した文字列を読む
=>  60:         l.readChar()
    61:         return tok
    62: }
    63:
    64: func newToken(tokenType token.TokenType, ch byte) token.Token {
    65:         return token.Token{Type: tokenType, Literal: string(ch)}
> github.com/pei0804/interpreter/lexer.(*Lexer).readChar() ./lexer.go:27 (PC: 0x10f194e)
Warning: listing may not match stale executable
    22:         return l
    23: }
    24:
    25: func (l *Lexer) readChar() {
    26:         // inputされた文字が最後まで参照したかをチェックする
                // まだ最後じゃないので終了しない
=>  27:         if l.readPosition >= len(l.input) {
    28:                 l.ch = 0
    29:         } else {
    30:                 l.ch = l.input[l.readPosition]
    31:         }
    32:         l.position = l.readPosition
(dlv) p len(l.input)
8
(dlv) p l
*github.com/pei0804/interpreter/lexer.Lexer {
        input: "=+(){},;",
        position: 0,
        readPosition: 1,
        ch: 61,}
> github.com/pei0804/interpreter/lexer.(*Lexer).NextToken() ./lexer.go:61 (PC: 0x10f1ad2)
Warning: listing may not match stale executable
    56:         case 0:
    57:                 tok.Literal = ""
    58:                 tok.Type = token.EOF
    59:         }
    60:         l.readChar()
=>  61:         return tok
    62: }
    63:
    64: func newToken(tokenType token.TokenType, ch byte) token.Token {
    65:         return token.Token{Type: tokenType, Literal: string(ch)}
    66: }
(dlv) p l
*github.com/pei0804/interpreter/lexer.Lexer {
        input: "=+(){},;",
        position: 1,
        readPosition: 2,
        ch: 43,}
> github.com/pei0804/interpreter/lexer.TestNextToken() ./lexer_test.go:39 (PC: 0x10f2421)
    34:                 if tok.Type != tt.expectedType {
    35:                         t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
    36:                                 i, tt.expectedType, tok.Type)
    37:                 }
    38:
=>  39:                 if tok.Literal != tt.expectedLiteral {
    40:                         t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
    41:                                 i, tt.expectedLiteral, tok.Literal)
    42:                 }
    43:         }
    44: }
(dlv) p tok
github.com/pei0804/interpreter/token.Token {Type: "=", Literal: "="}
(dlv) p tt
struct { github.com/pei0804/interpreter/lexer.expectedType github.com/pei0804/interpreter/token.TokenType; github.com/pei0804/inter
preter/lexer.expectedLiteral string } {expectedType: "=", expectedLiteral: "="}

// 最後の文字列
> github.com/pei0804/interpreter/lexer.(*Lexer).readChar() ./lexer.go:28 (PC: 0x10f1970)
Warning: listing may not match stale executable
    23: }
    24:
    25: func (l *Lexer) readChar() {
    26:         // inputされた文字が最後まで参照したかをチェックする
    27:         if l.readPosition >= len(l.input) {
                        // これ以上文字列が存在しないため0をchに入れる
=>  28:                 l.ch = 0
    29:         } else {
    30:                 l.ch = l.input[l.readPosition]
    31:         }
    32:         l.position = l.readPosition
    33:         l.readPosition += 1
(dlv) p l
*github.com/pei0804/interpreter/lexer.Lexer {
        input: "=+(){},;",
        position: 7,
        readPosition: 8,
        ch: 59,}
(dlv) p len(l.input)
8

> github.com/pei0804/interpreter/lexer.(*Lexer).readChar() ./lexer.go:34 (PC: 0x10f19b0)
Warning: listing may not match stale executable
    29:         } else {
    30:                 l.ch = l.input[l.readPosition]
    31:         }
    32:         l.position = l.readPosition
    33:         l.readPosition += 1
=>  34: }
    35:
    36: func (l *Lexer) NextToken() token.Token {
    37:         var tok token.Token
    38:
    39:         switch l.ch {
(dlv) p l
*github.com/pei0804/interpreter/lexer.Lexer {
        input: "=+(){},;",
        position: 8,
        readPosition: 9,
        ch: 0,}

// 終了を意味する0にヒットする
> github.com/pei0804/interpreter/lexer.(*Lexer).NextToken() ./lexer.go:56 (PC: 0x10f1a93)
Warning: listing may not match stale executable
    51:                 tok = newToken(token.PLUS, l.ch)
    52:         case '{':
    53:                 tok = newToken(token.LBRACE, l.ch)
    54:         case '}':
    55:                 tok = newToken(token.RBRACE, l.ch)
=>  56:         case 0:
    57:                 tok.Literal = ""
                        // EOFを入れる
    58:                 tok.Type = token.EOF
    59:         }
    60:         l.readChar()
    61:         return tok
```

