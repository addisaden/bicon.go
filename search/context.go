package search

import (
	"slices"
	"strings"

	"github.com/addisaden/bicon.go/base"
)

func SearchContext(context base.VerseTextType, selected []string, query string, limit int, offset int) ResultType {
	results := ResultType{query, limit, offset, context, selected, []ResultVerseType{}}
	query_split := strings.Fields(query)
	query_split_with_context := slices.Concat(selected, query_split)
	for i := range 66 {
		bookId := uint8(i + 1)
		thisBook := values(base.GetBooks(bookId))
		for _, verse := range thisBook {
			if len(results.Results) >= results.Limit {
				break
			}
			if versePresearch(verse, query_split_with_context) {
				verseAnalyzed := base.Analyze(verse)
				if selected_match, _ := verseSearchStrongnumber(verseAnalyzed, selected); !selected_match {
					continue
				}
				if search_matched, search_list := verseSearchComplete(verseAnalyzed, query_split); search_matched {
					results.Results = append(results.Results, ResultVerseType{verseAnalyzed, search_list})
				}
			}
		}
		if len(results.Results) >= results.Limit {
			break
		}
	}
	return results
}
