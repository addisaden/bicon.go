package base

import (
	"regexp"
)

func Analyze(input VerseType) VerseTextType {
	thisVerse := VerseTextType{
		input.Book,
		input.Chapter,
		input.Verse,
		input.Text,
		"",
		[]string{},
		false,
		[]VerseWordType{},
	}

	// Title
	title_area_reg := regexp.MustCompile(`<TS>([^<>]*).*<Ts>`)

	title_area := title_area_reg.FindStringSubmatch(input.Text)
	if len(title_area) >= 2 {
		thisVerse.Title = title_area[1]
	}

	// References
	ref_reg := regexp.MustCompile(`<RX\s+([^<>]+)>`)
	if len(title_area) >= 1 {
		refs := ref_reg.FindAllStringSubmatch(title_area[0], -1)
		for _, ref := range refs {
			if len(ref) >= 2 {
				thisVerse.Reference = append(thisVerse.Reference, ref[1])
			}
		}
	}

	// RTL
	rtl_reg := regexp.MustCompile(`<rtl\s*/>`)
	thisVerse.Rtl = rtl_reg.MatchString(input.Text)

	// Words
	words_reg := regexp.MustCompile(`<Q>.*?<q>`)

	for _, word_string := range words_reg.FindAllString(input.Text, -1) {
		this_word := VerseWordType{word_string, "", "", "", "", "", ""}
		// "<Q><H>בְּרֵאשִׁ֖ית<WH7225><WTPrep-b│N-fs><h><X>bə·rê·šîṯ<x><E> In the beginning <e><q>"
		// "<Q><G>Παῦλος<WG3972><WTN-NMS><g><X>Paulos<x><E> Paul <e><q>"
		hebrew_reg := regexp.MustCompile(`<H>.*<h>`)
		greek_reg := regexp.MustCompile(`<G>.*<g>`)

		if hebrew_reg.MatchString(word_string) {
			this_word.Language = "hebrew"
			text_reg := regexp.MustCompile(`<H>([^<>]+)<`)
			original_word := text_reg.FindStringSubmatch(word_string)
			if len(original_word) >= 2 {
				this_word.Text = original_word[1]
			}

			strong_reg := regexp.MustCompile(`<W(H\d+)>`)
			strong_match := strong_reg.FindStringSubmatch(word_string)
			if len(strong_match) >= 2 {
				this_word.Strongnumber = strong_match[1]
			}
		}
		if greek_reg.MatchString(word_string) {
			this_word.Language = "greek"
			text_reg := regexp.MustCompile(`<G>([^<>]+)<`)
			original_word := text_reg.FindStringSubmatch(word_string)
			if len(original_word) >= 2 {
				this_word.Text = original_word[1]
			}

			strong_reg := regexp.MustCompile(`<W(G\d+)>`)
			strong_match := strong_reg.FindStringSubmatch(word_string)
			if len(strong_match) >= 2 {
				this_word.Strongnumber = strong_match[1]
			}
		}

		grammar_reg := regexp.MustCompile(`<WT([^<>]*)>`)
		grammar_match := grammar_reg.FindStringSubmatch(word_string)
		if len(grammar_match) >= 2 {
			this_word.Grammar = grammar_match[1]
		}

		pronunciation_reg := regexp.MustCompile(`<X>(.*?)<x>`)
		pronunciation_match := pronunciation_reg.FindStringSubmatch(word_string)
		if len(pronunciation_match) >= 2 {
			this_word.Pronunciation = pronunciation_match[1]
		}

		english_reg := regexp.MustCompile(`<E>\s*(.*?)\s*<e>`)
		english_match := english_reg.FindStringSubmatch(word_string)
		if len(english_match) >= 2 {
			this_word.English = english_match[1]
		}

		thisVerse.Words = append(thisVerse.Words, this_word)
	}

	return thisVerse
}

func AnalyzeAll(inputmap map[string]VerseType) map[string]VerseTextType {
	result := map[string]VerseTextType{}

	for key, value := range inputmap {
		result[key] = Analyze(value)
	}

	return result
}
