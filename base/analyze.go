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

	// Words
	words_reg := regexp.MustCompile(`<Q>.*?<q>`)

	for _, word_string := range words_reg.FindAllString(input.text, -1) {
		this_word := VerseWordType{word_string, "", "", "", "", "", ""}
		// "<Q><H>בְּרֵאשִׁ֖ית<WH7225><WTPrep-b│N-fs><h><X>bə·rê·šîṯ<x><E> In the beginning <e><q>"
		// "<Q><G>Παῦλος<WG3972><WTN-NMS><g><X>Paulos<x><E> Paul <e><q>"
		hebrew_reg := regexp.MustCompile(`<H>.*<h>`)
		greek_reg := regexp.MustCompile(`<G>.*<g>`)

		if hebrew_reg.MatchString(word_string) {
			this_word.language = "hebrew"
			text_reg := regexp.MustCompile(`<H>([^<>]+)<`)
			original_word := text_reg.FindStringSubmatch(word_string)
			this_word.text = original_word[1]

			strong_reg := regexp.MustCompile(`<W(H\d+)>`)
			strong_match := strong_reg.FindStringSubmatch(word_string)
			this_word.strongnumber = strong_match[1]
		}
		if greek_reg.MatchString(word_string) {
			this_word.language = "greek"
			text_reg := regexp.MustCompile(`<G>([^<>]+)<`)
			original_word := text_reg.FindStringSubmatch(word_string)
			this_word.text = original_word[1]

			strong_reg := regexp.MustCompile(`<W(G\d+)>`)
			strong_match := strong_reg.FindStringSubmatch(word_string)
			this_word.strongnumber = strong_match[1]
		}

		grammar_reg := regexp.MustCompile(`<WT([^<>]*)>`)
		grammar_match := grammar_reg.FindStringSubmatch(word_string)
		this_word.grammar = grammar_match[1]

		pronunciation_reg := regexp.MustCompile(`<X>(.*?)<x>`)
		pronunciation_match := pronunciation_reg.FindStringSubmatch(word_string)
		this_word.pronunciation = pronunciation_match[1]

		english_reg := regexp.MustCompile(`<E>\s*(.*?)\s*<e>`)
		english_match := english_reg.FindStringSubmatch(word_string)
		this_word.english = english_match[1]

		thisVerse.words = append(thisVerse.words, this_word)
	}

	return thisVerse
}
