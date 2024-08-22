package base

import (
	"fmt"
	"testing"
)

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

	for i := range 66 {
		i_str := fmt.Sprint(i + 1)
		_, thisVerseFound := bereanInterlinear[i_str+"_1_1"]
		if !thisVerseFound {
			t.Errorf("%v. book is not found in the array bereanInterlinear", i_str)
		}
	}

	// t.Errorf("%v", firstVerse)
}
