package base

// To minify the overhead, another type is created
// Analyze(VerseType) => VerseTextType

type VerseTextType struct {
	Book      uint8
	Chapter   uint16
	Verse     uint16
	Text      string
	Title     string
	Reference []string
	Rtl       bool
	Words     []VerseWordType
}

type VerseWordType struct {
	Raw           string
	Language      string
	Text          string
	English       string
	Strongnumber  string
	Grammar       string
	Pronunciation string
}
