package keywords

type Keyword string

const (
	If        Keyword = "if"
	Else      Keyword = "else"
	For       Keyword = "for"
	Import    Keyword = "import"
	Component Keyword = "component"
)

var Keywords = []Keyword{
	If,
	Else,
	For,
	Import,
	Component,
}

// IsKeyword checks if the given word is a keyword.
func IsKeyword(word string) bool {
	kw := Keyword(word)
	for _, keyword := range Keywords {
		if keyword == kw {
			return true
		}
	}
	return false
}
