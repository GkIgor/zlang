package lexer

import (
	"strings"
)

// TODO implement lexer for the following grammar:
// image("path"), text("text"), attr(...attr)

// attr("key", "value")

// TODO implement a parser for the above grammar
// and then a lexer for the parser output

// TODO implement a tokenizer for the above grammar
// and then a lexer for the tokenizer output

// Define a token type
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

func Lex(input string) []Token {
	tokens := []Token{}
	lines := strings.Split(input, string(SEMICOLON))

	for _, line := range lines {
		line = strings.TrimSpace(line)
		// println(line)

		parts := strings.Split(line, string(ASSIGN))
		if len(parts) == 2 {

			LexVar(strings.Split(line, string(ASSIGN)), &tokens)

			// tokens = append(tokens, Token{Type: TokenVar, Value: strings.TrimSpace(parts[0])})

			// tokens = append(tokens, Token{Type: TokenValue, Value: strings.TrimSpace(parts[1])})

			if strings.TrimSpace(parts[1]) == string(TokenComponent) {

				// tokens = append(tokens, Token{Type: TokenComponent, Value: strings.TrimSpace(parts[1])})
			}

		}

		// else if len(strings.Split(line, string(ASSIGN))) != 2 {
		// 	fmt.Println(line)
		// 	panic("Syntax error on line: " + strconv.Itoa(index+1))
		// }

		// tokens = append(tokens, Token{Type: TokenLineEnd})
	}

	return tokens
}

// if len(line) == 0 || line == "\n" || line == "\r\n" {
// 	tokens = append(tokens, Token{Type: TokenLineEnd})
// 	continue
// }

// if line[len(line)-1] == ';' {
// 	line = line[:len(line)-1]
// }

// parts := strings.Split(line, ",")

// for _, part := range parts {
// 	part = strings.TrimSpace(part)

// 	if len(part) == 0 {
// 		continue
// 	}

// }
