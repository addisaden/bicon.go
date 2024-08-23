package base

import "regexp"

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

	// Title
	title_area_reg := regexp.MustCompile(`<TS>([^<>]*).*<Ts>`)

	title_area := title_area_reg.FindStringSubmatch(input.text)
	thisVerse.title = title_area[1]

	// References
	ref_reg := regexp.MustCompile(`<RX\s+([^<>]+)>`)
	refs := ref_reg.FindAllStringSubmatch(title_area[0], -1)
	for _, ref := range refs {
		thisVerse.reference = append(thisVerse.reference, ref[1])
	}

	// RTL
	rtl_reg := regexp.MustCompile(`<rtl\s*/>`)
	thisVerse.rtl = rtl_reg.MatchString(input.text)

	return thisVerse
}
