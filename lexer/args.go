package lexer

import "strings"

func LexParent(parts []string, tokens *[]Token) {
	if tokens == nil {
		panic("tokens cannot be nil")
	}

	if len(parts) < 2 {
		panic("Syntax error on line: " + getNonEmptyString(parts, 0))
	}

	trimmedLastPart := strings.TrimSpace(parts[len(parts)-1])
	trimmedFirstPart := strings.TrimSpace(parts[0])

	if trimmedLastPart == string(RPAREN) {
		bite := byte(strings.TrimSpace(parts[1])[0])

		var lexFunc func([]string, *[]Token)
		switch {
		case IsQuoteOrSingleQuote(bite):
			lexFunc = LexQuote
		case isDigit(bite):
			lexFunc = lexNumber
		case isLetter(bite):
			lexFunc = lexLetter
		case isSymbol(bite):
			lexFunc = lexRuneOrLetter
		default:
			panic("Syntax error on line: " + getNonEmptyString(parts, 0))
		}

		lexFunc(parts[1:len(parts)-1], tokens)
		*tokens = append(*tokens, Token{Type: RPAREN, Value: trimmedFirstPart})

	} else if trimmedFirstPart == string(LPAREN) {
		*tokens = append(*tokens, Token{Type: LPAREN, Value: trimmedFirstPart})
	} else {
		panic("Syntax error on line: " + getNonEmptyString(parts, 0))
	}

}
