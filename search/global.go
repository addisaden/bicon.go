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

func versePresearch(verse base.VerseType, queries []string) bool {
	all_queries_matched := len(queries)
	for _, q := range queries {
		if strings.Contains(strings.ToLower(verse.Text), strings.ToLower(q)) {
			all_queries_matched -= 1
		}
	}
	return all_queries_matched <= 0
}

func verseSearchStrongnumber(verseAnalyzed base.VerseTextType, queries []string) (bool, []int) {
	all_queries_matched := len(queries)
	matched_list := []int{}
	for _, q := range queries {
		matched := false
		for matched_index, word := range verseAnalyzed.Words {
			if strings.EqualFold(word.Strongnumber, q) {
				matched = true
				matched_list = append(matched_list, matched_index)
				break
			}
		}
		if matched {
			all_queries_matched -= 1
		}
	}
	return (all_queries_matched <= 1), matched_list
}

func verseSearchComplete(verseAnalyzed base.VerseTextType, queries []string) (bool, []int) {
	all_queries_matched := len(queries)
	matched_list := []int{}
	for _, q := range queries {
		matched := false
		for matched_index, word := range verseAnalyzed.Words {
			if strings.Contains(strings.ToLower(word.Text), strings.ToLower(q)) {
				matched = true
				matched_list = append(matched_list, matched_index)
				break
			}
			if strings.Contains(strings.ToLower(word.English), strings.ToLower(q)) {
				matched = true
				matched_list = append(matched_list, matched_index)
				break
			}
			if strings.EqualFold(word.Grammar, q) {
				matched = true
				matched_list = append(matched_list, matched_index)
				break
			}
			if strings.EqualFold(word.Strongnumber, q) {
				matched = true
				matched_list = append(matched_list, matched_index)
				break
			}
		}
		if matched {
			all_queries_matched -= 1
		}
	}
	return (all_queries_matched <= 1), matched_list
}

func SearchGlobal(query string, limit int, offset int) ResultType {
	results := ResultType{query, limit, offset, base.VerseTextType{}, []string{}, []ResultVerseType{}}
	query_split := strings.Fields(query)
	for i := range 66 {
		bookId := uint8(i + 1)
		thisBook := values(base.GetBooks(bookId))
		for _, verse := range thisBook {
			if len(results.Results) >= results.Limit {
				break
			}
			if versePresearch(verse, query_split) {
				verseAnalyzed := base.Analyze(verse)
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
