package base

import "testing"

func TestVerseTextType(t *testing.T) {
	genesis11 := "<TS>The Creation<br>(<RX 43.1.1-5> <RX 58.11.1-3>)<Ts><rtl/> <Q><H>בְּרֵאשִׁ֖ית<WH7225><WTPrep-b│N-fs><h><X>bə·rê·šîṯ<x><E> In the beginning <e><q> <Q><H>בָּרָ֣א<WH1254><WTV-Qal-Perf-3ms><h><X>bā·rā<x><E> created <e><q> <Q><H>אֱלֹהִ֑ים<WH430><WTN-mp><h><X>’ĕ·lō·hîm<x><E> God <e><q> <Q><H>אֵ֥ת<WH853><WTDirObjM><h><X>’êṯ<x><E> - <e><q> <Q><H>הַשָּׁמַ֖יִם<WH8064><WTArt│N-mp><h><X>haš·šā·ma·yim<x><E> the heavens <e><q> <Q><H>וְאֵ֥ת<WH853><WTConj-w│DirObjM><h><X>wə·’êṯ<x><E> and <e><q> <Q><H>הָאָֽרֶץ׃<WH776><WTArt│N-fs><h><X>hā·’ā·reṣ<x><E> the earth . <e><q>"

	nt5011text := "<TS>Greetings from Paul and Timothy<br>(<RX 51.1.1-2> <RX 57.1.1-3>)<Ts> <Q><G>Παῦλος<WG3972><WTN-NMS><g><X>Paulos<x><E> Paul <e><q> <Q><G>καὶ<WG2532><WTConj><g><X>kai<x><E> and <e><q> <Q><G>Τιμόθεος<WG5095><WTN-NMS><g><X>Timotheos<x><E> Timothy , <e><q> <Q><G>δοῦλοι<WG1401><WTN-NMP><g><X>douloi<x><E> servants <e><q> <Q><G>Χριστοῦ<WG5547><WTN-GMS><g><X>Christou<x><E> of Christ <e><q> <Q><G>Ἰησοῦ<WG2424><WTN-GMS><g><X>Iēsou<x><E> Jesus , <e><q> <Q><G>Πᾶσιν<WG3956><WTAdj-DMP><g><X>Pasin<x><E> To all <e><q> <Q><G>τοῖς<WG3588><WTArt-DMP><g><X>tois<x><E> the <e><q> <Q><G>ἁγίοις<WG40><WTAdj-DMP><g><X>hagiois<x><E> saints <e><q> <Q><G>ἐν<WG1722><WTPrep><g><X>en<x><E> in <e><q> <Q><G>Χριστῷ<WG5547><WTN-DMS><g><X>Christō<x><E> Christ <e><q> <Q><G>Ἰησοῦ<WG2424><WTN-DMS><g><X>Iēsou<x><E> Jesus <e><q> <Q><G>τοῖς<WG3588><WTArt-DMP><g><X>tois<x><E> - <e><q> <Q><G>οὖσιν<WG1510><WTV-PPA-DMP><g><X>ousin<x><E> - <e><q> <Q><G>ἐν<WG1722><WTPrep><g><X>en<x><E> at <e><q> <Q><G>Φιλίπποις<WG5375><WTN-DMP><g><X>Philippois<x><E> Philippi , <e><q> <Q><G>σὺν<WG4862><WTPrep><g><X>syn<x><E> together with <e><q> <Q><G>ἐπισκόποις<WG1985><WTN-DMP><g><X>episkopois<x><E> [the] overseers <e><q> <Q><G>καὶ<WG2532><WTConj><g><X>kai<x><E> and <e><q> <Q><G>διακόνοις<WG1249><WTN-DMP><g><X>diakonois<x><E> deacons : <e><q>"

	g11VerseType := VerseType{1, 1, 1, genesis11}
	ntVerseType := VerseType{50, 1, 1, nt5011text}

	g11 := Analyze(g11VerseType)
	nt5011 := Analyze(ntVerseType)

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

	if nt5011.title != "Greetings from Paul and Timothy" {
		t.Errorf("title %v must be 'Greetings from Paul and Timothy'", nt5011.title)
	}

	if len(g11.reference) != 2 {
		t.Errorf("there must be 2 references in %v", g11.reference)
	} else {
		if g11.reference[0] != "43.1.1-5" {
			t.Errorf("First reference %v must be '43.1.1-5'", g11.reference[0])
		}
	}

	if g11.rtl != true {
		t.Errorf("RTL should be true not %v", g11.rtl)
	}

	if nt5011.rtl != false {
		t.Errorf("RTL should be false not %v", nt5011.rtl)
	}

	if len(g11.words) != 7 {
		t.Errorf("It should be 7 words in %v", g11.words)
	} else {
		hfirst := g11.words[0]
		if hfirst.raw != "<Q><H>בְּרֵאשִׁ֖ית<WH7225><WTPrep-b│N-fs><h><X>bə·rê·šîṯ<x><E> In the beginning <e><q>" {
			t.Errorf("Raw is not present in %v", hfirst.raw)
		}
		if hfirst.text != "בְּרֵאשִׁ֖ית" {
			t.Errorf("Text %v is not present in %v", "בְּרֵאשִׁ֖ית", hfirst.text)
		}
		if hfirst.language != "hebrew" {
			t.Errorf("Language %v is not hebrew", hfirst.language)
		}
		if hfirst.strongnumber != "H7225" {
			t.Errorf("Strongnumber %v is not H7225", hfirst.strongnumber)
		}
		if hfirst.grammar != "Prep-b│N-fs" {
			t.Errorf("Grammar %v is not Prep-b│N-fs", hfirst.grammar)
		}
		if hfirst.pronunciation != "bə·rê·šîṯ" {
			t.Errorf("Pronunctiation %v is not bə·rê·šîṯ", hfirst.pronunciation)
		}
		if hfirst.english != "In the beginning" {
			t.Errorf("English %v is not 'In the beginning'", hfirst.english)
		}
	}

	if len(nt5011.words) != 20 {
		t.Errorf("It should be 20 words in %v", nt5011.words)
	} else {
		gfirst := nt5011.words[0]
		if gfirst.raw != "<Q><G>Παῦλος<WG3972><WTN-NMS><g><X>Paulos<x><E> Paul <e><q>" {
			t.Errorf("Raw is not present in %v", gfirst.raw)
		}
		if gfirst.text != "Παῦλος" {
			t.Errorf("Text %v is not present %v", "Παῦλος", gfirst.text)
		}
		if gfirst.language != "greek" {
			t.Errorf("Language %v is not greek", gfirst.language)
		}
		if gfirst.strongnumber != "G3972" {
			t.Errorf("Strongnumber %v is not G3972", gfirst.strongnumber)
		}
		if gfirst.grammar != "N-NMS" {
			t.Errorf("Grammar %v is not N-NMS", gfirst.grammar)
		}
		if gfirst.pronunciation != "Paulos" {
			t.Errorf("Pronunctiation %v is not Paulos", gfirst.pronunciation)
		}
		if gfirst.english != "Paul" {
			t.Errorf("English %v is not 'Paul'", gfirst.english)
		}
	}
}
