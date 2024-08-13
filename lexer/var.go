package lexer

import "strings"

// Function to lex variables

func LexVar(parts []string, tokens *[]Token) {
	*tokens = append(*tokens, Token{Type: TokenVar, Value: strings.TrimSpace(parts[0])})

	*tokens = append(*tokens, Token{Type: TokenValue, Value: strings.TrimSpace(parts[1])})
}
