package lexer

import (
	"reflect"
	"testing"
)

func TestLex(t *testing.T) {

	tests := []struct {
		name   string
		input  string
		output []Token
	}{
		{
			name:   "Empty input",
			input:  "",
			output: []Token{},
		},
		{
			name:   "Single line with image token",
			input:  "image(\"path\");",
			output: []Token{{Type: TokenImage, Value: "path"}, {Type: TokenLineEnd}},
		},
		{
			name:   "Single line with text token",
			input:  "text(\"text\");",
			output: []Token{{Type: TokenText, Value: "text"}, {Type: TokenLineEnd}},
		},
		{
			name:   "Single line with multiple attribute tokens",
			input:  "attr(\"key1\", \"value1\", \"key2\", \"value2\");",
			output: []Token{{Type: TokenKey, Value: "key1"}, {Type: TokenValue, Value: "value1"}, {Type: TokenKey, Value: "key2"}, {Type: TokenValue, Value: "value2"}, {Type: TokenLineEnd}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Lex(tt.input)

			if !reflect.DeepEqual(got, tt.output) {
				t.Errorf("Lex(%s) = %v, want %v", tt.input, got, tt.output)
			}
		})
	}
}
