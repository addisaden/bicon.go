package base

import "testing"

func TestBereanInterlinear(t *testing.T) {
	bereanInterlinear := GetBereanInterlinear()

	_, firstVerseFound := bereanInterlinear["1_1_1"]
	if !firstVerseFound {
		t.Error("Error: First verse is empty")
	}

	_, lastVerseFound := bereanInterlinear["66_22_21"]
	if !lastVerseFound {
		t.Error("Error: Last verse is empty")
	}

}
