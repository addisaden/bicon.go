package search

import (
	"fmt"
	"testing"
)

func TestGlobalSimple(t *testing.T) {
	searchTerm1 := "H7225"
	maxResults1 := 1
	offset1 := 0
	verseIndex1 := "1_1_1"

	searchResults := SearchGlobal(searchTerm1, maxResults1, offset1)
	if searchResults.Query != searchTerm1 {
		t.Errorf("Query %v must be %v", searchResults.Query, searchTerm1)
	}
	if searchResults.Limit != maxResults1 {
		t.Errorf("Limit %v must be %v", searchResults.Limit, maxResults1)
	}
	if searchResults.Offset != offset1 {
		t.Errorf("Offset %v must be %v", searchResults.Offset, offset1)
	}
	if len(searchResults.Results) != 1 {
		t.Errorf("len(Results) %v must be %v", len(searchResults.Results), 1)
	}
	firstResult := searchResults.Results[0]
	verseIndex := fmt.Sprintf("%d_%d_%d", firstResult.Book, firstResult.Chapter, firstResult.Verse)
	if verseIndex != verseIndex1 {
		t.Errorf("First match %v must be %v", verseIndex, verseIndex1)
	}
}

func TestGlobalMultiTerm(t *testing.T) {
	searchTerm := "H1254 H7225" // the 2nd and the 1st word from the first verse
	searchResults := SearchGlobal(searchTerm, 1, 0)
	if len(searchResults.Results) < 1 {
		t.Errorf("len(Results) is %v must be 1", len(searchResults.Results))
	} else {
		firstResult := searchResults.Results[0]
		verseIndex := fmt.Sprintf("%d_%d_%d", firstResult.Book, firstResult.Chapter, firstResult.Verse)
		if verseIndex != "1_1_1" {
			t.Errorf("First match %v must be %v", verseIndex, "1_1_1")
		}
	}
}

func TestGlobalNoResultsPerformance(t *testing.T) {
	searchTerm := "asfjoangewgsngndfionqwoinwfgioanfioanfoiwqngiowq"
	searchResults := SearchGlobal(searchTerm, 1, 0)
	if len(searchResults.Results) > 0 {
		t.Errorf("len(Results) is %v must be 0", len(searchResults.Results))
	}
}
