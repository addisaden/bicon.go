package base

import "testing"

func TestVerseTextType(t *testing.T) {
	genesis11 := "<TS>The Creation<br>(<RX 43.1.1-5> <RX 58.11.1-3>)<Ts><rtl/> <Q><H>בְּרֵאשִׁ֖ית<WH7225><WTPrep-b│N-fs><h><X>bə·rê·šîṯ<x><E> In the beginning <e><q> <Q><H>בָּרָ֣א<WH1254><WTV-Qal-Perf-3ms><h><X>bā·rā<x><E> created <e><q> <Q><H>אֱלֹהִ֑ים<WH430><WTN-mp><h><X>’ĕ·lō·hîm<x><E> God <e><q> <Q><H>אֵ֥ת<WH853><WTDirObjM><h><X>’êṯ<x><E> - <e><q> <Q><H>הַשָּׁמַ֖יִם<WH8064><WTArt│N-mp><h><X>haš·šā·ma·yim<x><E> the heavens <e><q> <Q><H>וְאֵ֥ת<WH853><WTConj-w│DirObjM><h><X>wə·’êṯ<x><E> and <e><q> <Q><H>הָאָֽרֶץ׃<WH776><WTArt│N-fs><h><X>hā·’ā·reṣ<x><E> the earth . <e><q>"

	g11VerseType := VerseType{1, 1, 1, genesis11}

	g11 := Analyze(g11VerseType)

	if g11.book != 1 {
		t.Errorf("book %v must be 1", g11.book)
	}

	if g11.chapter != 1 {
		t.Errorf("book %v must be 1", g11.chapter)
	}

	if g11.verse != 1 {
		t.Errorf("book %v must be 1", g11.verse)
	}

	if g11.text != genesis11 {
		t.Errorf("text\n%v\nmust be\n%v", g11.text, genesis11)
	}

	if g11.title != "The Creation" {
		t.Errorf("title %v must be 'The Creation'", g11.title)
	}

	if len(g11.reference) != 2 {
		t.Errorf("there must be 2 references in %v", g11.reference)
	} else {
		if g11.reference[0] != "43.1.1-5" {
			t.Errorf("First reference %v must be '43.1.1-5'", g11.reference[0])
		}
	}
}
