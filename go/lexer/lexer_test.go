package lexer 

import(
  "testing"
  "go/token"
)

func TestNextToken(t *testing.T){
  input := `=+(){},;`

  tests := []struct{
    expectedType token.TokenType
    expectedLiteral string
  }{
    {token.ASSIGN, "="},
    {token.PLUS, "+"},
    {token.LPARAN, "("},
    {token.RPAREN, ")"},
    {token.LBRACE, "{"},
    {token.RBRACE, "}"},
    {token.COMMA, ","},
    {token.EOF, ""},
    {token.SEMICOLON, ";"},
  }

  l := New(input)
  for i, tt := range tests{
    tok := l.NextToken()
    if tok.type != tt.expectedType{
      t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%d", i, tt.expectedType, tok.type)
    }
  }
}
