package lexer

import "strings"

// LexQuote parses a line containing a quoted string
func LexQuote(parts []string, tokens *[]Token) {
	if tokens == nil {
		panic("tokens cannot be nil")
	}

	if len(parts) < 2 {
		panic("Syntax error on line: " + getNonEmptyString(parts, 0))
	}

	startQuote := parts[0][0]
	endQuote := parts[len(parts)-1][len(parts[len(parts)-1])-1]

	if startQuote != endQuote || !IsQuoteOrSingleQuote(startQuote) {
		panic("Syntax error on line: " + getNonEmptyString(parts, 0))
	}

	trimmedStartQuote := strings.TrimSpace(parts[0])
	trimmedEndQuote := strings.TrimSpace(parts[len(parts)-1])

	*tokens = append(*tokens, Token{Type: QUOTE, Value: trimmedStartQuote})
	*tokens = append(*tokens, Token{Type: STRING, Value: trimAndJoinParts(parts)})
	*tokens = append(*tokens, Token{Type: QUOTE, Value: trimmedEndQuote})
}

func getNonEmptyString(parts []string, index int) string {
	if index < len(parts) && parts[index] != "" {
		return parts[index]
	}
	return ""
}

func trimAndJoinParts(parts []string) string {
	trimmedParts := make([]string, len(parts)-1)
	for i, part := range parts[:len(parts)-1] {
		trimmedParts[i] = strings.TrimSpace(part)
	}
	return strings.Join(trimmedParts, "")
}

func IsQuoteOrSingleQuote(ch byte) bool {
	return ch == '"' || ch == '\''
}
