package search

import (
	"cmp"
	"slices"
	"strings"

	"github.com/addisaden/bicon.go/base"
)

func values(input map[string]base.VerseType) []base.VerseType {
	results := []base.VerseType{}
	for _, value := range input {
		results = append(results, value)
	}
	slices.SortFunc(results, func(a, b base.VerseType) int {
		if a.Book != b.Book {
			return cmp.Compare(a.Book, b.Book)
		}
		if a.Chapter != b.Chapter {
			return cmp.Compare(a.Chapter, b.Chapter)
		}
		return cmp.Compare(a.Verse, b.Verse)
	})
	return results
}

func SearchGlobal(query string, limit int, offset int) ResultType {
	results := ResultType{query, limit, offset, []base.VerseTextType{}}
	query_split := strings.Fields(query)
	for i := range 66 {
		bookId := uint8(i + 1)
		thisBook := values(base.GetBooks(bookId))
		for _, verse := range thisBook {
			if len(results.Results) >= results.Limit {
				break
			}
			verseAnalyzed := base.Analyze(verse)
			all_query_matched := len(query_split)
			for _, q := range query_split {
				matched := false
				for _, word := range verseAnalyzed.Words {
					if strings.Contains(strings.ToLower(word.Text), strings.ToLower(q)) {
						matched = true
						break
					}
					if strings.Contains(strings.ToLower(word.English), strings.ToLower(q)) {
						matched = true
						break
					}
					if strings.EqualFold(word.Grammar, q) {
						matched = true
						break
					}
					if strings.EqualFold(word.Strongnumber, q) {
						matched = true
						break
					}
				}
				if matched {
					all_query_matched -= 1
				}
			}
			if all_query_matched <= 0 {
				results.Results = append(results.Results, verseAnalyzed)
				break
			}
		}
		if len(results.Results) >= results.Limit {
			break
		}
	}
	return results
}
