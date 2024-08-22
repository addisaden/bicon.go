package base

func Analyze(input VerseType) VerseTextType {
	thisVerse := VerseTextType{
		input.book,
		input.chapter,
		input.verse,
		input.text,
		"",
		[]string{},
		false,
		[]VerseWordType{},
	}
	return thisVerse
}
