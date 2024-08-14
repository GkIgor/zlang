package lexer

import "strings"

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isSymbol(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') || ch == '_' || ch == '-' || ch == '.' || ch == ':'
}

func lexRune(parts []string, tokens *[]Token) {
	*tokens = append(*tokens, Token{Type: TokenValue, Value: strings.TrimSpace(parts[0])})
}

func lexLetter(parts []string, tokens *[]Token) {
	*tokens = append(*tokens, Token{Type: TokenValue, Value: strings.TrimSpace(parts[0])})
}

func lexNumber(parts []string, tokens *[]Token) {
	*tokens = append(*tokens, Token{Type: TokenValue, Value: strings.TrimSpace(parts[0])})
}

func lexRuneOrLetter(parts []string, tokens *[]Token) {
	*tokens = append(*tokens, Token{Type: TokenValue, Value: strings.TrimSpace(parts[0])})
}

func lexRuneOrNumber(parts []string, tokens *[]Token) {
	*tokens = append(*tokens, Token{Type: TokenValue, Value: strings.TrimSpace(parts[0])})
}
