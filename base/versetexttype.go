package base

// To minify the overhead, another type is created
// Analyze(VerseType) => VerseTextType

type VerseTextType struct {
	book      uint8
	chapter   uint16
	verse     uint16
	text      string
	title     string
	reference []string
	rtl       bool
	words     []VerseWordType
}

type VerseWordType struct {
	raw           string
	language      string
	text          string
	english       string
	strongnumber  string
	grammar       string
	pronunciation string
}
