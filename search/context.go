package search

import (
	"github.com/addisaden/bicon.go/base"
)

func SearchContext(context base.VerseTextType, selected []string, query string, limit int, offset int) ResultType {
	results := ResultType{query, limit, offset, context, selected, []ResultVerseType{}}
	return results
}
