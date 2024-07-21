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
	TokenImage   TokenType = "image"
	TokenText    TokenType = "text"
	TokenAttr    TokenType = "attr"
	TokenKey     TokenType = "key"
	TokenValue   TokenType = "value"
	TokenLineEnd TokenType = "lineEnd"
)

// Define a token structure
type Token struct {
	Type  TokenType
	Value string
}

// Implement the lexer
func Lex(input string) []Token {
	tokens := []Token{}
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		if line[len(line)-1] == ';' {
			line = line[:len(line)-1]
		}

		parts := strings.Split(line, ",")

		for _, part := range parts {
			part = strings.TrimSpace(part)

			if strings.HasPrefix(part, "image") {
				tokens = append(tokens, Token{Type: TokenImage, Value: part[5 : len(part)-1]})
			} else if strings.HasPrefix(part, "text") {
				tokens = append(tokens, Token{Type: TokenText, Value: part[5 : len(part)-1]})
			} else if strings.HasPrefix(part, "attr") {
				attr := part[5 : len(part)-1]
				attrParts := strings.Split(attr, "\" ")

				for i := 0; i < len(attrParts); i += 2 {
					tokens = append(tokens, Token{Type: TokenKey, Value: attrParts[i][1 : len(attrParts[i])-1]})
					tokens = append(tokens, Token{Type: TokenValue, Value: attrParts[i+1][1 : len(attrParts[i+1])-1]})
				}
			}
		}

		tokens = append(tokens, Token{Type: TokenLineEnd})
	}

	return tokens
}
