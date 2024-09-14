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
		[]string{"H430", "H8064", "H776"}, // These strongs must be present (if empty at least one)
		"wife",                            // search term -> 1_24_3
		3,                                 // Search Results
		0,                                 // Offset
	)
	t.Errorf("%v", results)
}
