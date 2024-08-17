package base

import "testing"

func TestVerseType(t *testing.T) {
	genesis11 := "<TS>The Creation<br>(<RX 43.1.1-5> <RX 58.11.1-3>)<Ts><rtl/> <Q><H>בְּרֵאשִׁ֖ית<WH7225><WTPrep-b│N-fs><h><X>bə·rê·šîṯ<x><E> In the beginning <e><q> <Q><H>בָּרָ֣א<WH1254><WTV-Qal-Perf-3ms><h><X>bā·rā<x><E> created <e><q> <Q><H>אֱלֹהִ֑ים<WH430><WTN-mp><h><X>’ĕ·lō·hîm<x><E> God <e><q> <Q><H>אֵ֥ת<WH853><WTDirObjM><h><X>’êṯ<x><E> - <e><q> <Q><H>הַשָּׁמַ֖יִם<WH8064><WTArt│N-mp><h><X>haš·šā·ma·yim<x><E> the heavens <e><q> <Q><H>וְאֵ֥ת<WH853><WTConj-w│DirObjM><h><X>wə·’êṯ<x><E> and <e><q> <Q><H>הָאָֽרֶץ׃<WH776><WTArt│N-fs><h><X>hā·’ā·reṣ<x><E> the earth . <e><q>"

	// testVerse := VerseType{1, 1, 1, genesis11}

	testVerseList := []VerseType{
		{1, 1, 1, genesis11},
	}

	testVerse := testVerseList[0]

	if testVerse.book != 1 {
		t.Errorf("testVerse.book: %v is not 1: %v", testVerse.book, testVerse)
	}
	if testVerse.chapter != 1 {
		t.Errorf("testVerse.chapter: %v is not 1: %v", testVerse.chapter, testVerse)
	}
	if testVerse.verse != 1 {
		t.Errorf("testVerse.verse: %v is not 1: %v", testVerse.verse, testVerse)
	}
	if testVerse.text != genesis11 {
		t.Errorf("testVerse.verse: %v is not %v: %v", testVerse.text, genesis11, testVerse)
	}
}
