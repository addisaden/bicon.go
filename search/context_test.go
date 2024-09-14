package search

import (
	"testing"

	"github.com/addisaden/bicon.go/base"
)

func TestContextSimple(t *testing.T) {
	contextVerse := "1_1_1"
	genesis := base.GetBereanInterlinear1()
	context := base.Analyze(genesis[contextVerse])

	results := SearchContext(
		context,                           // At least one strong (or selection) must be present in the match
		[]string{"H430", "H8064", "H776"}, // These strongs must be present
		"wife",                            // search term -> 1_24_3
		3,                                 // Search Results
		0,                                 // Offset
	)
	firstResult := results.Results[0].Verse
	if !(firstResult.Book == 1 && firstResult.Chapter == 24 && firstResult.Verse == 3) {
		t.Errorf("The first result must be 1_24_3, but is %v", firstResult)
	}
	for _, result := range results.Results {
		if matched, _ := verseSearchComplete(result.Verse, []string{"H430", "H8064", "H776", "wife"}); !matched {
			t.Errorf("Something is missing in this verse %v", result)
		}
	}
}
