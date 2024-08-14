package lexer

type TokenType string

const (
	TokenImage     TokenType = "image"
	TokenText      TokenType = "text"
	TokenKey       TokenType = "key"
	TokenValue     TokenType = "value"
	TokenVar       TokenType = "VAR"
	TokenComponent TokenType = "component"
	TokenLineEnd   TokenType = "lineEnd"
	STRING         TokenType = "STRING"
	NUMBER         TokenType = "NUMBER"
	IDENT          TokenType = "IDENT"
	EOF            TokenType = "EOF"
	ILLEGAL        TokenType = "ILLEGAL"
	COMMA          TokenType = ","
	LPAREN         TokenType = "("
	RPAREN         TokenType = ")"
	SEMICOLON      TokenType = ";"
	ASSIGN         TokenType = "="
	NEWLINE        TokenType = "\n"
	QUOTE          TokenType = "\""
)

// Define a token structure
type Token struct {
	Type  TokenType
	Value string
}

// type Lexer struct {
// 	input        string
// 	position     int
// 	readPosition int
// 	ch           byte
// }
