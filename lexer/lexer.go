package lexer

import (
	"fmt"
	"strconv"
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
	TokenImage   TokenType = "image"
	TokenText    TokenType = "text"
	TokenKey     TokenType = "key"
	TokenValue   TokenType = "value"
	TokenLineEnd TokenType = "lineEnd"
	ILLEGAL      TokenType = "ILLEGAL"
	EOF          TokenType = "EOF"
	IDENT        TokenType = "IDENT"
	COMMA        TokenType = ","
	LPAREN       TokenType = "("
	RPAREN       TokenType = ")"
	STRING       TokenType = "STRING"
	NUMBER       TokenType = "NUMBER"
	SEMICOLON    TokenType = ";"
	ASSIGN       TokenType = "="
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

	for index, line := range lines {
		line = strings.TrimSpace(line)

		// println(line)

		if len(strings.Split(line, string(ASSIGN))) == 2 {
			fmt.Println(strings.Split(line, string(ASSIGN)))

		} else if len(strings.Split(line, string(ASSIGN))) != 2 {
			fmt.Println(line)
			panic("Syntax error on line: " + strconv.Itoa(index+1))
		}

		tokens = append(tokens, Token{Type: TokenLineEnd})
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
