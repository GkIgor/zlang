package keywords

type Keyword string

const (
	IF        Keyword = "if"
	ELSE      Keyword = "else"
	FOR       Keyword = "for"
	IMPORT    Keyword = "import"
	COMPONENT Keyword = "component"
)

var keywords = []Keyword{
	IF,
	ELSE,
	FOR,
	IMPORT,
	COMPONENT,
}

func IsKeyword(word string) bool {
	for _, keyword := range keywords {
		if string(keyword) == word {
			return true
		}
	}
	return false
}
