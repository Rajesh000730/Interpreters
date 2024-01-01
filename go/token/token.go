package token 

type TokenType string

type token struct{
  Type TokenType
  Literal string
}

const (
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"

  //identifiers + literals
  IDENT = "IDENT"
  INT = "INT"

  //Operators
  PLUS = "+"
  ASSIGN = "="

  //delemiters
  COMMA = ","
  SEMICOLON = ";"

  LPARAN = "("
  RPARAN = ")"
  LBRACE = "{"
  RBRACE = "}"

  //keywords
  FUNCTION = "FUNCTION"
  LET = "LET"
)
