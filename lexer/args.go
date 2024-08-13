package lexer

import "strings"

func LexParent(parts []string, tokens *[]Token) {

	if strings.TrimSpace(parts[0]) == string(RPAREN) {

		bite := byte(strings.TrimSpace(parts[1])[0])

		if IsQuoteOrSingleQuote(bite) {

		}

		*tokens = append(*tokens, Token{Type: RPAREN, Value: strings.TrimSpace(parts[0])})
	}

	if strings.TrimSpace(parts[0]) == string(LPAREN) {

		*tokens = append(*tokens, Token{Type: LPAREN, Value: strings.TrimSpace(parts[0])})
	}

}
